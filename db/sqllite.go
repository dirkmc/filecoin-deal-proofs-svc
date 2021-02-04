package db

import (
	"crypto/sha256"
	"database/sql"
	"time"

	"github.com/cbergoon/merkletree"
	"github.com/davecgh/go-spew/spew"
	logging "github.com/ipfs/go-log/v2"
)

const dsn = "./mock/sqlite-database.db"

var log = logging.Logger("db")

type DB interface {
	DealByID(dealID uint64) (*Deal, error)
	GetAllDeals() error
	Close() error
}

type liteDB struct {
	conn *sql.DB
}

func New() (DB, error) {
	conn, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	return &liteDB{conn: conn}, nil
}

func (db *liteDB) Close() error {
	return db.Close()
}

type Deal struct {
	DealID      uint64
	DataCID     string
	PieceCID    string
	Provider    string
	StartEpoch  uint64
	EndEpoch    uint64
	SignedEpoch uint64
}

func (d Deal) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(d.Provider)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func (d Deal) Equals(other merkletree.Content) (bool, error) {
	return d.Provider == other.(Deal).Provider, nil
}

type Deals []*Deal

func (ds Deals) Root() string {
	return "0x1234"
}

func (db *liteDB) GetAllDeals() error {
	// drop all deals from db

	log.Debugw("dropping deals db")
	statement, _ := db.conn.Prepare(`DROP TABLE deals`)
	statement.Exec()

	// create db based on data from sentinel

	createDb := `CREATE TABLE deals (
		"dealId" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"dataCid" TEXT,
		"pieceCid" TEXT,
		"provider" TEXT,
		"startEpoch" integer,
		"endEpoch" integer,
		"signedEpoch" integer
	  );`

	log.Debugw("create `deals` table")
	statement, err := db.conn.Prepare(createDb)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Debugw("`deals` table created")

	now := time.Now()

	var deals []merkletree.Content

	for i := 1; i < 10; i++ {
		d := &Deal{
			DealID:      uint64(now.UnixNano()),
			DataCID:     "datacid1234",
			PieceCID:    "piececid1234",
			Provider:    "fprovider1",
			StartEpoch:  10,
			EndEpoch:    2000,
			SignedEpoch: 50,
		}

		log.Debugw("generated deal", "deal", spew.Sdump(d))

		deals = append(deals, d)
	}

	tree, err := merkletree.NewTree(deals)
	if err != nil {
		panic(err)
	}

	mr := tree.MerkleRoot()
	log.Debugw("merkle root", "root", mr)

	return nil

}

func (db *liteDB) DealByID(dealID uint64) (*Deal, error) {
	statement, err := db.conn.Prepare("SELECT * FROM deals WHERE DealID = ?")
	if err != nil {
		return nil, err
	}

	return RowToDeal(statement.QueryRow(dealID))
}

type Scannable interface {
	Scan(dest ...interface{}) error
}

func RowToDeal(row Scannable) (*Deal, error) {
	var deal Deal
	err := row.Scan(
		&deal.DealID,
		&deal.DataCID,
		&deal.PieceCID,
		&deal.Provider,
		&deal.StartEpoch,
		&deal.EndEpoch,
		&deal.SignedEpoch)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &deal, nil
}
