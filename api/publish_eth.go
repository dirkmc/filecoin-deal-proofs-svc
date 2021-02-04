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
	endpoint = "https://rinkeby.infura.io/v3/7e1eddb52ae149eaaa92941def0fd49d"
	contract = common.HexToAddress("0xd4375467f6CfB0493b5e4AF0601B3a0f2e7D2FcA")

	managerAddress = common.HexToAddress("0x3b8Fd7cE0f4841F1C23B67b20676886ac230Be64")
	privateKey     = "f0ce4b609fe0865dd37595908c2c01e5e8ca887983f6db638f5ffe5b3067887c"

	signedEpoch = big.NewInt(50)
)

func publishMerkleRootToEthereum() {
	if db.MerkleRoot == "" {
		panic("empty merkle root")
	}

	client, err := ethclient.Dial(endpoint)
	if err != nil {
		panic(err)
	}

	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic(err)
	}

	signFn := func(s types.Signer, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
		rinkeby := big.NewInt(4)
		return types.SignTx(tx, types.NewEIP155Signer(rinkeby), pk)
	}

	opts := &bind.TransactOpts{
		From:     managerAddress,
		GasPrice: big.NewInt(5000000000), // 5 gwei
		Signer:   signFn,
		Nonce:    nil,
		Value:    nil,
		GasLimit: uint64(1500000),
	}

	fo, err := oracle.NewFilecoinService(contract, client)
	if err != nil {
		panic(err)
	}

	var mr [32]byte
	copy(mr[:], common.HexToAddress(db.MerkleRoot).Bytes())

	tx, err := fo.UpdateState(opts, mr, signedEpoch)
	if err != nil {
		panic(err)
	}

	log.Debugw("sent tx", "txhash", tx.Hash())
}
