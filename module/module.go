package module

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/hyperledger-labs/yui-relayer/config"
	"github.com/hyperledger-labs/yui-relayer/core"
	"github.com/spf13/cobra"
)

type Module struct{}

var _ config.ModuleI = (*Module)(nil)

// Name returns the name of the module
func (Module) Name() string {
	return "quorum"
}

// RegisterInterfaces register the module interfaces to protobuf Any.
func (Module) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*core.ChainConfigI)(nil),
		&ChainConfig{},
	)
	registry.RegisterImplementations(
		(*core.ProverConfigI)(nil),
		&ProverConfig{},
	)
}

// GetCmd returns the command
func (Module) GetCmd(ctx *config.Context) *cobra.Command {
	return nil
}
