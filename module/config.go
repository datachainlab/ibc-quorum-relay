package module

import (
	fmt "fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hyperledger-labs/yui-relayer/core"
)

var _ core.ChainConfigI = (*ChainConfig)(nil)
var _ core.ProverConfigI = (*ProverConfig)(nil)

func (c ChainConfig) Build() (core.ChainI, error) {
	return NewChain(c)
}

func (c ProverConfig) Build(chain core.ChainI) (core.ProverI, error) {
	chain_, ok := chain.(*Chain)
	if !ok {
		return nil, fmt.Errorf("chain type must be %T, not %T", &Chain{}, chain)
	}
	return NewProver(chain_, c), nil
}

func (c ChainConfig) IBCHostAddress() common.Address {
	return common.HexToAddress(c.IbcHostAddress)
}

func (c ChainConfig) IBCHandlerAddress() common.Address {
	return common.HexToAddress(c.IbcHandlerAddress)
}
