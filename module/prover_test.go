package module_test

import (
	fmt "fmt"
	"testing"

	mocktypes "github.com/datachainlab/ibc-mock-client/modules/light-clients/xx-mock/types"
	"github.com/datachainlab/ibc-quorum-relay/module"
	"github.com/hyperledger-labs/yui-relayer/core"
)

const (
	// any private key is ok because quorum tx is fee-free
	privateKey = "0x0000000000000000000000000000000000000000000000000000000000000001"
	// contract address changes for each deployment
	ibcHandlerAddress = "0x6468751F5D94540338058254D8F9BD1AcEa498Fe"
)

func makeChain() (*module.Chain, error) {
	// instantiate a chain module
	chain, err := module.NewChain(module.ChainConfig{
		RpcAddr:           "http://localhost:8545",
		EthChainId:        1337,
		PrivateKey:        privateKey,
		IbcHandlerAddress: ibcHandlerAddress,
	})
	if err != nil {
		return nil, fmt.Errorf("NewChain failed: %w", err)
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

func TestChain(t *testing.T) {
	chain, err := makeChain()
	if err != nil {
		t.Fatalf("makeChain failed: %v", err)
	}

	bn, err := chain.GetLatestHeight()
	if err != nil {
		t.Fatalf("chain.GetLatestHeight failed: %v", err)
	}
	if _, err := chain.QueryClientState(bn); err != nil {
		t.Fatalf("chain.QueryClientState failed: %v", err)
	}
	if _, err := chain.QueryConnection(bn); err != nil {
		t.Fatalf("chain.QueryConnection failed: %v", err)
	}
	if _, err := chain.QueryChannel(bn); err != nil {
		t.Fatalf("chain.QueryChannel failed: %v", err)
	}
}

func TestProver(t *testing.T) {
	chain, err := makeChain()
	if err != nil {
		t.Fatalf("makeChain failed: %v", err)
	}
	prover := module.NewProver(chain, module.ProverConfig{
		TrustLevelNumerator:   0,
		TrustLevelDenominator: 0,
		TrustingPeriod:        0,
	})
	if _, err := prover.QueryLatestHeader(); err != nil {
		t.Fatalf("prover.QueryLatestHeader failed: %v", err)
	}
	bn, err := prover.GetLatestLightHeight()
	if err != nil {
		t.Fatalf("prover.GetLatestLightHeight failed: %v", err)
	}
	if _, err := prover.QueryClientStateWithProof(bn); err != nil {
		t.Fatalf("prover.QueryClientStateWithProof failed: %v", err)
	}
	if _, err := prover.QueryConnectionWithProof(bn); err != nil {
		t.Fatalf("prover.QueryConnectionWithProof failed: %v", err)
	}
	if _, err := prover.QueryChannelWithProof(bn); err != nil {
		t.Fatalf("prover.QueryChannelWithProof failed: %v", err)
	}
}
