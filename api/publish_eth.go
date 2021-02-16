package api

import (
	"math/big"

	"github.com/dirkmc/filecoin-deal-proofs-svc/bindings/oracle"
	"github.com/dirkmc/filecoin-deal-proofs-svc/db"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	signedEpoch = big.NewInt(50)

	Production bool
	Prvkey     string

	ChainId        *big.Int
	Endpoint       string
	OracleContract common.Address
	ManagerAddress common.Address
)

func publishMerkleRootToEthereum() {
	if db.MerkleRoot == "" {
		panic("empty merkle root")
	}

	client, err := ethclient.Dial(Endpoint)
	if err != nil {
		panic(err)
	}

	pk, err := crypto.HexToECDSA(Prvkey)
	if err != nil {
		panic(err)
	}

	signFn := func(s types.Signer, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(tx, types.NewEIP155Signer(ChainId), pk)
	}

	opts := &bind.TransactOpts{
		From:     ManagerAddress,
		GasPrice: big.NewInt(5000000000), // 5 gwei
		Signer:   signFn,
		Nonce:    nil,
		Value:    nil,
		GasLimit: uint64(1500000),
	}

	fo, err := oracle.NewFilecoinService(OracleContract, client)
	if err != nil {
		panic(err)
	}

	// TODO: cheat for merkle root
	db.MerkleRoot = "0xfbf13f17c0c2c5b0e2f79c31578632d0f44f07b92aa078d8d44e5252cbdbba1b"

	var mr [32]byte
	copy(mr[:], common.HexToHash(db.MerkleRoot).Bytes())

	if Production {
		tx, err := fo.UpdateState(opts, mr, signedEpoch)
		if err != nil {
			panic(err)
		}

		log.Infow("sent tx", "txhash", tx.Hash())
	} else {
		log.Debugw("running in debug mode, so not sending tx to ethereum with merkle root")
	}
}
