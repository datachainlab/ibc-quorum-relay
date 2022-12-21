package module

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchandler"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchost"
	"github.com/hyperledger-labs/yui-relayer/core"
)

type Chain struct {
	config ChainConfig

	rpcClient  *rpc.Client
	ethClient  *ethclient.Client
	ethChainID *big.Int
	privateKey *ecdsa.PrivateKey
	ibcHost    *ibchost.Ibchost
	ibcHandler *ibchandler.Ibchandler

	homePath         string
	codec            codec.ProtoCodecMarshaler
	pathEnd          *core.PathEnd
	msgEventListener core.MsgEventListener
}

var _ core.ChainI = (*Chain)(nil)

func NewChain(config ChainConfig) (*Chain, error) {
	rpcClient, err := rpc.Dial(config.RpcAddr)
	if err != nil {
		return nil, err
	}

	ethClient := ethclient.NewClient(rpcClient)

	ethChainID := big.NewInt(config.EthChainId)

	hexPrivKey := config.PrivateKey
	if hexPrivKey[:2] == "0x" {
		hexPrivKey = hexPrivKey[2:]
	}
	privateKey, err := crypto.HexToECDSA(hexPrivKey)
	if err != nil {
		return nil, err
	}

	ibcHost, err := ibchost.NewIbchost(config.IBCHostAddress(), ethClient)
	if err != nil {
		return nil, err
	}

	ibcHandler, err := ibchandler.NewIbchandler(config.IBCHandlerAddress(), ethClient)
	if err != nil {
		return nil, err
	}

	return &Chain{
		config:     config,
		rpcClient:  rpcClient,
		ethClient:  ethClient,
		ethChainID: ethChainID,
		privateKey: privateKey,
		ibcHost:    ibcHost,
		ibcHandler: ibcHandler,
	}, nil
}

// ChainID returns ID of the chain
func (c *Chain) ChainID() string {
	return c.config.ChainId
}

// GetLatestHeight gets the chain for the latest height and returns it
func (c *Chain) GetLatestHeight() (int64, error) {
	bn, err := c.ethClient.BlockNumber(context.TODO())
	if err != nil {
		return 0, err
	}
	return int64(bn), nil
}

// GetAddress returns the address of relayer
func (c *Chain) GetAddress() (sdk.AccAddress, error) {
	ethAddr := crypto.PubkeyToAddress(c.privateKey.PublicKey)
	return sdk.AccAddress(ethAddr[:]), nil
}

// Codec returns the codec
func (c *Chain) Codec() codec.ProtoCodecMarshaler {
	return c.codec
}

// SetRelayInfo sets source's path and counterparty's info to the chain
func (c *Chain) SetRelayInfo(p *core.PathEnd, counterparty *core.ProvableChain, counterpartyPath *core.PathEnd) error {
	if err := p.Validate(); err != nil {
		return err
	}
	c.pathEnd = p
	return nil
}

// Path returns the path
func (c *Chain) Path() *core.PathEnd {
	return c.pathEnd
}

// Init initializes the chain
func (c *Chain) Init(homePath string, timeout time.Duration, codec codec.ProtoCodecMarshaler, debug bool) error {
	c.homePath = homePath
	c.codec = codec
	return nil
}

// SetupForRelay performs chain-specific setup before starting the relay
func (c *Chain) SetupForRelay(ctx context.Context) error {
	return nil
}

// RegisterMsgEventListener registers a given EventListener to the chain
func (c *Chain) RegisterMsgEventListener(listener core.MsgEventListener) {
	c.msgEventListener = listener
}
