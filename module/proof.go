package module

import (
	fmt "fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

type accountResult struct {
	Address      common.Address  `json:"address"`
	AccountProof []string        `json:"accountProof"`
	Balance      *hexutil.Big    `json:"balance"`
	CodeHash     common.Hash     `json:"codeHash"`
	Nonce        hexutil.Uint64  `json:"nonce"`
	StorageHash  common.Hash     `json:"storageHash"`
	StorageProof []storageResult `json:"storageProof"`
}
type storageResult struct {
	Key   string       `json:"key"`
	Value *hexutil.Big `json:"value"`
	Proof []string     `json:"proof"`
}

func (pr *Prover) ethGetProof(address common.Address, storageKeys []common.Hash, height int64) (*accountResult, error) {
	blockNumber := "0x" + strconv.FormatInt(height, 16)
	var accountResult accountResult
	if err := pr.chain.rpcClient.Call(&accountResult, "eth_getProof", address, storageKeys, blockNumber); err != nil {
		return nil, err
	}
	return &accountResult, nil
}

func (pr *Prover) getAccountProof(height int64) ([]byte, error) {
	// call eth_getProof
	accountResult, err := pr.ethGetProof(
		pr.chain.config.IBCHostAddress(),
		nil,
		height,
	)
	if err != nil {
		return nil, err
	}

	// get account proof
	var accountProof [][]byte
	for _, hexProof := range accountResult.AccountProof {
		proof, err := hexutil.Decode(hexProof)
		if err != nil {
			return nil, err
		}
		accountProof = append(accountProof, proof)
	}

	// encode account proof
	rlpAccountProof, err := rlp.EncodeToBytes(accountProof)
	if err != nil {
		return nil, err
	}

	return rlpAccountProof, nil
}

func (pr *Prover) getStateCommitmentProof(path []byte, height int64) ([]byte, error) {
	// calculate slot for commitment
	slot := crypto.Keccak256Hash(append(
		crypto.Keccak256Hash(path).Bytes(),
		common.Hash{}.Bytes()...,
	))

	// call eth_getProof
	accountResult, err := pr.ethGetProof(
		pr.chain.config.IBCHostAddress(),
		[]common.Hash{slot},
		height,
	)
	if err != nil {
		return nil, err
	}

	// get storage proof
	var storageProof [][]byte
	for _, hexProof := range accountResult.StorageProof[0].Proof {
		proof, err := hexutil.Decode(hexProof)
		if err != nil {
			return nil, err
		}
		storageProof = append(storageProof, proof)
	}

	// encode storage proof
	rlpStorageProof, err := rlp.EncodeToBytes(storageProof)
	if err != nil {
		return nil, err
	}

	return rlpStorageProof, nil
}

type proofList struct {
	list  [][]byte
	index int
}

func (p *proofList) Has([]byte) (bool, error) {
	panic("not implemented")
}

func (p *proofList) Get([]byte) ([]byte, error) {
	if p.index >= len(p.list) {
		return nil, fmt.Errorf("out of index")
	}
	v := p.list[p.index]
	p.index += 1
	return v, nil
}

func verifyProof(rootHash common.Hash, key []byte, proof [][]byte) ([]byte, error) {
	return trie.VerifyProof(rootHash, key, &proofList{list: proof, index: 0})
}
