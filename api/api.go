package api

import (
	"encoding/json"
	"fmt"
	"github.com/dirkmc/filecoin-deal-proofs-svc/db"
	logging "github.com/ipfs/go-log/v2"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/xerrors"
	"net/http"
	"strconv"
)

var log = logging.Logger("daemon")

type API struct {
	db db.DB
}

func New(db db.DB) *API {
	return &API{db: db}
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

	w.Write(json)
}

func sendErrResponse(w http.ResponseWriter, err error) {
	json := []byte(fmt.Sprintf(`{"Error":"%s"}`, err))
	w.Write(json)
}
