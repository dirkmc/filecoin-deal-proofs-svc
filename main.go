package main

import (
	"context"
	"github.com/dirkmc/filecoin-deal-proofs-svc/db"
	"github.com/dirkmc/filecoin-deal-proofs-svc/web"
	logging "github.com/ipfs/go-log/v2"
	"net/http"
	"time"

	"github.com/dirkmc/filecoin-deal-proofs-svc/api"
)

var log = logging.Logger("svc")

func main() {
	logging.SetAllLoggers(logging.LevelDebug)

	err := run()
	if err != nil {
		log.Error(err)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(ProcessContext())
	defer cancel()

	//cfg := &config.EnvConfig{}
	//if err := cfg.Load(); err != nil {
	//	return err
	//}
	//
	//srv, err := daemon.New(cfg)

	apidb, err := db.New()
	if err != nil {
		return err
	}

	srv, err := web.New(api.New(apidb))
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