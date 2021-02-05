package main

import (
	"flag"
	"math/big"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/davecgh/go-spew/spew"
	"github.com/dirkmc/filecoin-deal-proofs-svc/bindings/oracle"
	"github.com/dirkmc/filecoin-deal-proofs-svc/db"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("verify-proof")

var (
	OurAddress = common.HexToAddress("0x3b8Fd7cE0f4841F1C23B67b20676886ac230Be64")

	endpoint = "https://rinkeby.infura.io/v3/7e1eddb52ae149eaaa92941def0fd49d"
	contract = common.HexToAddress("0xd4375467f6CfB0493b5e4AF0601B3a0f2e7D2FcA")

	dealID      int
	dataCID     string
	pieceCID    string
	provider    string
	startEpoch  int
	endEpoch    int
	signedEpoch int
	proof       string
)

func init() {
	flag.IntVar(&dealID, "dealID", 0, "")
	flag.StringVar(&dataCID, "dataCID", "", "")
	flag.StringVar(&pieceCID, "pieceCID", "", "")
	flag.StringVar(&provider, "provider", "", "")
	flag.IntVar(&startEpoch, "startEpoch", 0, "")
	flag.IntVar(&endEpoch, "endEpoch", 0, "")
	flag.IntVar(&signedEpoch, "signedEpoch", 0, "")
	flag.StringVar(&proof, "proof", "", "")
}

func main() {
	flag.Parse()

	client, err := rpc.Dial(endpoint)
	if err != nil {
		panic(err)
	}
	ethClient := ethclient.NewClient(client)

	fs, err := oracle.NewFilecoinService(contract, ethClient)
	if err != nil {
		panic(err)
	}

	d := &db.Deal{
		DealID:      big.NewInt(int64(dealID)),
		DataCID:     dataCID,
		PieceCID:    pieceCID,
		Provider:    provider,
		StartEpoch:  big.NewInt(int64(startEpoch)),
		EndEpoch:    big.NewInt(int64(endEpoch)),
		SignedEpoch: big.NewInt(int64(signedEpoch)),
	}

	spew.Dump(d)

	var merkleProof [][32]byte

	spew.Dump(len(proof))

	entries := len(proof) / 66

	spew.Dump(entries)
	for i := 0; i < entries; i++ {
		start := i * 66
		end := (i + 1) * 66

		entry := proof[start:end]
		spew.Dump(entry)

		slice := common.HexToHash(entry).Bytes()
		var arr [32]byte

		copy(arr[:], slice)

		merkleProof = append(merkleProof, arr)
	}

	spew.Dump(merkleProof)

	tx, err := fs.VerifyProof(nil, d.DataCID, d.PieceCID, d.DealID, d.Provider, d.StartEpoch, d.EndEpoch, d.SignedEpoch, merkleProof)
	if err != nil {
		panic(err)
	}

	log.Info("verifying a correct proof:")
	spew.Dump(tx)

	log.Info("verifying a broken proof:")
	spew.Dump(tx)
}
