package main

import (
	"context"
	"flag"
	"math/big"
	"net/http"
	"time"

	"github.com/dirkmc/filecoin-deal-proofs-svc/db"
	"github.com/dirkmc/filecoin-deal-proofs-svc/web"
	"github.com/ethereum/go-ethereum/common"
	logging "github.com/ipfs/go-log/v2"

	"github.com/dirkmc/filecoin-deal-proofs-svc/api"
)

var log = logging.Logger("svc")

var (
	production = flag.Bool("production", false, "run in prod, and send tx to ethereum rinkeby")
	prvkey     = flag.String("prvkey", "", "private key of account")
	remotedb   = flag.String("remotedb", "", "remote database")
	endpoint   = flag.String("endpoint", "https://rinkeby.infura.io/v3/xxxxx", "endpoint to an ethereum node")
	oracle     = flag.String("oracle", "0xd4375467f6CfB0493b5e4AF0601B3a0f2e7D2FcA", "oracle contract address on ethereum")
	manager    = flag.String("manager", "0x3b8Fd7cE0f4841F1C23B67b20676886ac230Be64", "manager address for the oracle contract")
	chainId    = flag.Int("chainid", 4, "chain id; rinkeby == 4")
)

func main() {
	logging.SetAllLoggers(logging.LevelDebug)

	flag.Parse()

	api.Production = *production
	api.Prvkey = *prvkey
	api.Endpoint = *endpoint
	api.OracleContract = common.HexToAddress(*oracle)
	api.ManagerAddress = common.HexToAddress(*manager)
	api.ChainId = big.NewInt(int64(*chainId))

	err := run()
	if err != nil {
		log.Error(err)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(ProcessContext())
	defer cancel()

	apidb, err := db.New(*remotedb)
	if err != nil {
		return err
	}

	a := api.New(apidb)
	a.FetchDealsPeriodically()

	srv, err := web.New(a)
	if err != nil {
		return err
	}

	exiting := make(chan struct{})
	defer close(exiting)

	go func() {
		select {
		case <-ctx.Done():
		case <-exiting:
			// no need to shutdown in this case.
			return
		}

		log.Infow("shutting down rpc server")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalw("failed to shut down rpc server", "err", err)
		}
		log.Infow("rpc server stopped")
	}()

	log.Infow("listen and serve", "addr", srv.Addr())
	err = srv.Serve()
	if err == http.ErrServerClosed {
		err = nil
	}
	return err
}
