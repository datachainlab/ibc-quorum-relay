package module

import (
	"log"

	"github.com/cosmos/ibc-go/v4/modules/core/exported"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/hyperledger-labs/yui-relayer/core"
)

var _ core.HeaderI = (*Header)(nil)

const Quorum string = "xx-quorum"

func (*Header) ClientType() string {
	return Quorum
}

func (h *Header) GetHeight() exported.Height {
	ethHeader, err := h.decodeEthHeader()
	if err != nil {
		log.Panicf("invalid header: %v", h)
	}
	return ethHeightToPB(ethHeader.Number.Int64())
}

func (h *Header) ValidateBasic() error {
	if _, err := h.decodeEthHeader(); err != nil {
		return err
	}
	if _, err := h.decodeAccountProof(); err != nil {
		return err
	}
	return nil
}

func (h *Header) decodeEthHeader() (*types.Header, error) {
	var ethHeader types.Header
	if err := rlp.DecodeBytes(h.GoQuorumHeaderRlp, &ethHeader); err != nil {
		return nil, err
	}
	return &ethHeader, nil
}

func (h *Header) decodeAccountProof() ([][]byte, error) {
	var accountProof [][]byte
	if err := rlp.DecodeBytes(h.AccountProof, &accountProof); err != nil {
		return nil, err
	}
	return accountProof, nil
}
