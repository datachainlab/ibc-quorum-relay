package module

import (
	"bytes"
	fmt "fmt"
	"testing"

	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	mocktypes "github.com/datachainlab/ibc-mock-client/modules/light-clients/xx-mock/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	proto "github.com/gogo/protobuf/proto"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/relay/ethereum"
	"github.com/hyperledger-labs/yui-relayer/core"
)

const (
	hdwMnemonic = "math razor capable expose worth grape metal sunset metal sudden usage scheme"
	hdwPath     = "m/44'/60'/0'/0/0"

	// contract address changes for each deployment
	ibcHandlerAddress = "0x702E40245797c5a2108A566b3CE2Bf14Bc6aF841"
)

func makeChain() (*ethereum.Chain, error) {
	// instantiate a chain module
	chain, err := ethereum.NewChain(ethereum.ChainConfig{
		RpcAddr:           "http://localhost:8545",
		EthChainId:        1337,
		HdwMnemonic:       hdwMnemonic,
		HdwPath:           hdwPath,
		IbcHandlerAddress: ibcHandlerAddress,
	})
	if err != nil {
		return nil, fmt.Errorf("ethereum.NewChain failed: %w", err)
	}

	// call Init
	codec := core.MakeCodec()
	mocktypes.RegisterInterfaces(codec.InterfaceRegistry())
	if err := chain.Init("", 0, codec, false); err != nil {
		return nil, fmt.Errorf("chain.Init failed: %w", err)
	}

	// call SetRelayInfo
	if err := chain.SetRelayInfo(&core.PathEnd{
		ClientID:     "mock-client-0",
		ConnectionID: "connection-0",
		ChannelID:    "channel-0",
		PortID:       "transfer",
		Order:        "UNORDERED",
	}, nil, nil); err != nil {
		return nil, fmt.Errorf("chain.SetRelayInfo failed: %w", err)
	}

	return chain, nil
}

func TestProver(t *testing.T) {
	// instantiate a prover module
	chain, err := makeChain()
	if err != nil {
		t.Fatalf("makeChain failed: %v", err)
	}
	prover := NewProver(chain, ProverConfig{
		TrustLevelNumerator:   0,
		TrustLevelDenominator: 0,
		TrustingPeriod:        0,
	})

	// get height used for testing
	bn, err := prover.GetLatestLightHeight()
	if err != nil {
		t.Fatalf("prover.GetLatestLightHeight failed: %v", err)
	}

	// test header queries
	if _, err := prover.QueryLatestHeader(); err != nil {
		t.Errorf("prover.QueryLatestHeader failed: %v", err)
	}
	iHeader, err := prover.QueryHeader(bn)
	if err != nil {
		t.Errorf("prover.QueryHeader failed: %v", err)
	}
	header := iHeader.(*Header)
	contractAddress := prover.chain.Config().IBCHandlerAddress()
	storageRoot := verifyHeader(t, header, contractAddress)

	// test client and consensus queries
	if res, err := prover.QueryClientStateWithProof(bn); err != nil {
		t.Errorf("prover.QueryClientStateWithProof failed: %v", err)
	} else {
		path := host.FullClientStatePath(prover.chain.Path().ClientID)
		commitment := messageToCommitment(t, res.ClientState)
		verifyMembership(t, storageRoot, res.Proof, path, commitment)
		if cs, err := clienttypes.UnpackClientState(res.ClientState); err != nil {
			t.Errorf("clienttypes.UnpackClientState failed: %v", err)
		} else if res, err := prover.QueryClientConsensusStateWithProof(bn, cs.GetLatestHeight()); err != nil {
			t.Errorf("prover.QueryClientConsensusStateWithProof failed: %v", err)
		} else {
			path := host.FullConsensusStatePath(prover.chain.Path().ClientID, cs.GetLatestHeight())
			commitment := messageToCommitment(t, res.ConsensusState)
			verifyMembership(t, storageRoot, res.Proof, path, commitment)
		}
	}

	// test connection query
	if res, err := prover.QueryConnectionWithProof(bn); err != nil {
		t.Errorf("prover.QueryConnectionWithProof failed: %v", err)
	} else {
		path := host.ConnectionPath(prover.chain.Path().ConnectionID)
		commitment := messageToCommitment(t, res.Connection)
		verifyMembership(t, storageRoot, res.Proof, path, commitment)
	}

	// test channel query
	if res, err := prover.QueryChannelWithProof(bn); err != nil {
		t.Errorf("prover.QueryChannelWithProof failed: %v", err)
	} else {
		path := host.ChannelPath(prover.chain.Path().PortID, prover.chain.Path().ChannelID)
		commitment := messageToCommitment(t, res.Channel)
		verifyMembership(t, storageRoot, res.Proof, path, commitment)
	}

	// test packet commitment query
	if res, err := prover.QueryPacketCommitmentWithProof(bn, 1); err != nil {
		t.Errorf("prover.QueryPacketCommitmentWithProof failed: %v", err)
	} else {
		path := host.PacketCommitmentPath(
			prover.chain.Path().PortID,
			prover.chain.Path().ChannelID,
			1,
		)
		verifyMembership(t, storageRoot, res.Proof, path, res.Commitment)
	}

	// test ack commitment query
	if res, err := prover.QueryPacketAcknowledgementCommitmentWithProof(bn, 1); err != nil {
		t.Errorf("prover.QueryPacketAcknowledgementCommitmentWithProof failed: %v", err)
	} else {
		path := host.PacketAcknowledgementPath(
			prover.chain.Path().PortID,
			prover.chain.Path().ChannelID,
			1,
		)
		verifyMembership(t, storageRoot, res.Proof, path, res.Acknowledgement)
	}
}

func messageToCommitment(t *testing.T, msg proto.Message) []byte {
	t.Helper()
	marshaled, err := proto.Marshal(msg)
	if err != nil {
		t.Fatalf("proto.Marshal(msg) failed: %v", err)
	}
	commitment := crypto.Keccak256(marshaled)
	return commitment
}

func verifyHeader(t *testing.T, header *Header, contractAddress common.Address) common.Hash {
	t.Helper()
	var rawAccountProof [][][]byte
	if err := rlp.DecodeBytes(header.AccountProof, &rawAccountProof); err != nil {
		t.Fatalf("rlp.DecodeBytes(header.AccountProof, ...) failed: %v", err)
	}
	var accountProof [][]byte
	for _, raw := range rawAccountProof {
		bz, err := rlp.EncodeToBytes(raw)
		if err != nil {
			t.Fatalf("rlp.EncodeToBytes(raw) failed: %v", err)
		}
		accountProof = append(accountProof, bz)
	}

	var quorumHeader types.Header
	if err := rlp.DecodeBytes(header.GoQuorumHeaderRlp, &quorumHeader); err != nil {
		t.Fatalf("rlp.DecodeBytes(header.GoQuorumHeaderRLP, ...) failed: %v", err)
	}

	var account state.Account
	if rlpAccount, err := verifyProof(
		quorumHeader.Root,
		crypto.Keccak256(contractAddress.Bytes()),
		accountProof,
	); err != nil {
		t.Fatalf("verifyProof failed: %v", err)
	} else if err := rlp.DecodeBytes(rlpAccount, &account); err != nil {
		t.Fatalf("rlp.DecodeBytes(rlpAccount, ...) failed: %v", err)
	}

	return account.Root
}

func verifyMembership(t *testing.T, root common.Hash, bzValueProof []byte, path string, commitment []byte) {
	t.Helper()
	var rawValueProof [][][]byte
	if err := rlp.DecodeBytes(bzValueProof, &rawValueProof); err != nil {
		t.Fatalf("rlp.DecodeBytes(bzValueProof, ...) failed: %v", err)
	}
	var valueProof [][]byte
	for _, raw := range rawValueProof {
		if bz, err := rlp.EncodeToBytes(raw); err != nil {
			t.Fatalf("rlp.EncodeToBytes(raw) failed: %v", err)
		} else {
			valueProof = append(valueProof, bz)
		}
	}

	key := crypto.Keccak256(crypto.Keccak256(append(crypto.Keccak256([]byte(path)), common.Hash{}.Bytes()...)))

	recoveredCommitment, err := verifyProof(root, key, valueProof)
	if err != nil {
		t.Fatalf("verifyProof failed: %v", err)
	}

	rlpCommitment, err := rlp.EncodeToBytes(commitment)
	if err != nil {
		t.Fatalf("rlp.EncodeToBytes(commitment) failed: %v", err)
	}
	if !bytes.Equal(recoveredCommitment, rlpCommitment) {
		t.Fatalf("value unmatch: %v(length=%d) != %v(length=%d)",
			recoveredCommitment, len(recoveredCommitment),
			rlpCommitment, len(rlpCommitment),
		)
	}
}
