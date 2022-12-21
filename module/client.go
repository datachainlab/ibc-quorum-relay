package module

import (
	"context"
	fmt "fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func (chain *Chain) CallOpts(ctx context.Context, height int64) *bind.CallOpts {
	opts := &bind.CallOpts{
		From:    crypto.PubkeyToAddress(chain.privateKey.PublicKey),
		Context: ctx,
	}
	if height > 0 {
		opts.BlockNumber = big.NewInt(height)
	}
	return opts
}

func (chain *Chain) TransactOpts(ctx context.Context) *bind.TransactOpts {
	signer := types.NewEIP155Signer(chain.ethChainID)
	privKey := chain.privateKey
	addr := crypto.PubkeyToAddress(privKey.PublicKey)
	return &bind.TransactOpts{
		From:     addr,
		GasLimit: 6382056,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != addr {
				return nil, fmt.Errorf("unexpected address: %s", address)
			}
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), privKey)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}
}
