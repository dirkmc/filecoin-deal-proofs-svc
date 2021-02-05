package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/dirkmc/filecoin-deal-proofs-svc/db"
	"github.com/dirkmc/filecoin-deal-proofs-svc/web"
	logging "github.com/ipfs/go-log/v2"

	"github.com/dirkmc/filecoin-deal-proofs-svc/api"
)

var log = logging.Logger("svc")

var production = flag.Bool("production", false, "run in prod, and send tx to ethereum rinkeby")

func main() {
	logging.SetAllLoggers(logging.LevelDebug)

	flag.Parse()

	api.Production = *production

	err := run()
	if err != nil {
		log.Error(err)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(ProcessContext())
	defer cancel()

	apidb, err := db.New(os.Getenv("REMOTE_DB"))
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
