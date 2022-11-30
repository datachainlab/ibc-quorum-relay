package module

import (
	"errors"

	"github.com/hyperledger-labs/yui-relayer/core"
)

var _ core.ChainConfigI = (*ChainConfig)(nil)
var _ core.ProverConfigI = (*ProverConfig)(nil)

func (c ChainConfig) Build() (core.ChainI, error) {
	return nil, errors.New("not implemented yet")
}

func (c ProverConfig) Build(chain core.ChainI) (core.ProverI, error) {
	return nil, errors.New("not implemented yet")
}
