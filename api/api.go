package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/dirkmc/filecoin-deal-proofs-svc/db"
	logging "github.com/ipfs/go-log/v2"
	"golang.org/x/xerrors"
)

var log = logging.Logger("daemon")

type API struct {
	dbmu sync.Mutex
	db   db.DB
}

func New(db db.DB) *API {
	return &API{db: db}
}

func (d *API) FetchDealsPeriodically() {
	//ticker := time.NewTicker(3600 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan struct{})

	// TODO: this channel is needed for demo to block the ticker after first iteration
	smth := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				return
			case _ = <-ticker.C:
				log.Debugw("fetch deals")

				d.fetchDeals()

				<-smth

			}
		}
	}()
}

func (d *API) fetchDeals() {
	d.dbmu.Lock()
	err := d.db.GetAllDeals()
	d.dbmu.Unlock()

	if err != nil {
		log.Errorf("Error fetching deals: %s", err)
		return
	}

	publishMerkleRootToEthereum()
}

func (d *API) DealHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := d.dealById(w, r)
		sendResponse(w, r, res, err)
	}
}

type NotFound struct {
	Message string
}

func (d *API) dealById(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	dealIDStr := r.URL.Query().Get("dealID")
	if dealIDStr == "" {
		return nil, xerrors.Errorf("missing query param `dealID`")
	}
	dealID, err := strconv.Atoi(dealIDStr)
	if err != nil {
		return nil, xerrors.Errorf("parsing query param `dealID`: %w", err)
	}
	deal, err := d.db.DealByID(uint64(dealID))
	if deal == nil {
		return NotFound{Message: fmt.Sprintf("deal with deal id %d not found", dealID)}, nil
	}
	deal.MerkleRoot = db.MerkleRoot // TODO: remove global merkle root
	return deal, err
}

func sendResponse(w http.ResponseWriter, r *http.Request, res interface{}, err error) {
	rlog := log.With("req_id", r.Header.Get("X-Request-ID"))
	rlog.Debugw("handle request", "action", "deal")
	defer rlog.Debugw("request handled", "action", "deal")

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		sendErrResponse(w, err)
		return
	}

	json, err := json.Marshal(res)
	if err != nil {
		sendErrResponse(w, err)
		return
	}

	//enableCors(&w)

	_, err = w.Write(json)
	if err != nil {
		log.Errorf("sending response: %s", err)
	}
}

func sendErrResponse(w http.ResponseWriter, err error) {
	json := []byte(fmt.Sprintf(`{"Error":"%s"}`, err))
	_, err = w.Write(json)
	if err != nil {
		log.Errorf("sending response: %s", err)
	}
}
