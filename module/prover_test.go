package module_test

import (
	fmt "fmt"
	"testing"

	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	mocktypes "github.com/datachainlab/ibc-mock-client/modules/light-clients/xx-mock/types"
	"github.com/datachainlab/ibc-quorum-relay/module"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/relay/ethereum"
	"github.com/hyperledger-labs/yui-relayer/core"
)

const (
	hdwMnemonic = "math razor capable expose worth grape metal sunset metal sudden usage scheme"
	hdwPath     = "m/44'/60'/0'/0/0"

	// contract address changes for each deployment
	ibcHandlerAddress = "0xaa43d337145E8930d01cb4E60Abf6595C692921E"
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
	prover := module.NewProver(chain, module.ProverConfig{
		TrustLevelNumerator:   0,
		TrustLevelDenominator: 0,
		TrustingPeriod:        0,
	})

	// get height used for testing
	bn, err := prover.GetLatestLightHeight()
	if err != nil {
		t.Fatalf("prover.GetLatestLightHeight failed: %v", err)
	}

	// test queries
	if _, err := prover.QueryLatestHeader(); err != nil {
		t.Errorf("prover.QueryLatestHeader failed: %v", err)
	}
	if csRes, err := prover.QueryClientStateWithProof(bn); err != nil {
		t.Errorf("prover.QueryClientStateWithProof failed: %v", err)
	} else if cs, err := clienttypes.UnpackClientState(csRes.ClientState); err != nil {
		t.Errorf("clienttypes.UnpackClientState failed: %v", err)
	} else if _, err := prover.QueryClientConsensusStateWithProof(bn, cs.GetLatestHeight()); err != nil {
		t.Errorf("prover.QueryClientConsensusStateWithProof failed: %v", err)
	}
	if _, err := prover.QueryConnectionWithProof(bn); err != nil {
		t.Errorf("prover.QueryConnectionWithProof failed: %v", err)
	}
	if _, err := prover.QueryChannelWithProof(bn); err != nil {
		t.Errorf("prover.QueryChannelWithProof failed: %v", err)
	}
	if _, err := prover.QueryPacketCommitmentWithProof(bn, 1); err != nil {
		t.Errorf("prover.QueryPacketCommitmentWithProof failed: %v", err)
	}
	if _, err := prover.QueryPacketAcknowledgementCommitmentWithProof(bn, 1); err != nil {
		t.Errorf("prover.QueryPacketAcknowledgementCommitmentWithProof failed: %v", err)
	}
}
