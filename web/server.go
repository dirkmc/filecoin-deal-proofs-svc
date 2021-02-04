package web

import (
	"context"
	"fmt"
	"github.com/dirkmc/filecoin-deal-proofs-svc/api"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	logging "github.com/ipfs/go-log/v2"
	"github.com/pborman/uuid"
)

var log = logging.Logger("daemon")

type Server struct {
	server *http.Server
	l      net.Listener
	doneCh chan struct{}
}

func New(api *api.API) (srv *Server, err error) {
	srv = new(Server)

	r := mux.NewRouter().StrictSlash(true)

	// Set a unique request ID.
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Set("X-Request-ID", uuid.New()[:8])
			next.ServeHTTP(w, r)
		})
	})

	staticDir := "/static/"
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	r.HandleFunc("/deal", api.DealHandler()).Methods("GET")

	srv.doneCh = make(chan struct{})
	srv.server = &http.Server{
		Handler:      r,
		WriteTimeout: 7200 * time.Second,
		ReadTimeout:  7200 * time.Second,
	}

	srv.l, err = net.Listen("tcp", "127.0.0.1:9518")
	if err != nil {
		return nil, err
	}

	return srv, nil
}

// Serve starts the server and blocks until the server is closed, either
// explicitly via Shutdown, or due to a fault condition. It propagates the
// non-nil err return value from http.Serve.
func (d *Server) Serve() error {
	select {
	case <-d.doneCh:
		return fmt.Errorf("tried to reuse a stopped server")
	default:
	}

	log.Infow("daemon listening", "addr", d.Addr())
	return d.server.Serve(d.l)
}

func (d *Server) Addr() string {
	return d.l.Addr().String()
}

func (d *Server) Port() int {
	return d.l.Addr().(*net.TCPAddr).Port
}

func (d *Server) Shutdown(ctx context.Context) error {
	defer close(d.doneCh)
	return d.server.Shutdown(ctx)
}

