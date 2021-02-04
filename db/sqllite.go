package db

import (
	"bytes"
	"database/sql"
	"math/big"
	"time"

	"github.com/cbergoon/merkletree"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	logging "github.com/ipfs/go-log/v2"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

const dsn = "./mock/sqlite-database.db"

var log = logging.Logger("db")
var MerkleRoot string

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
	MerkleRoot  string
	DealID      *big.Int
	DataCID     string
	PieceCID    string
	Provider    string
	StartEpoch  *big.Int
	EndEpoch    *big.Int
	SignedEpoch *big.Int
	Proof       string
}

func (d *Deal) CalculateHash() ([]byte, error) {
	types := []string{"string", "string", "uint256", "string", "uint256", "uint256", "uint256"}

	values := []interface{}{
		d.DataCID,
		d.PieceCID,
		d.DealID,
		d.Provider,
		d.StartEpoch,
		d.EndEpoch,
		d.SignedEpoch,
	}

	return solsha3.SoliditySHA3(types, values), nil
}

func (d *Deal) Equals(other merkletree.Content) (bool, error) {
	_self, _ := d.CalculateHash()
	_other, _ := other.CalculateHash()

	return bytes.Equal(_self, _other), nil
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
		"signedEpoch" integer,
		"proof" TEXT
	  );`

	log.Debugw("create `deals` table")
	statement, err := db.conn.Prepare(createDb)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Debugw("`deals` table created")

	//now := time.Now()

	// make deal ids deterministic
	shortForm := "2006-Jan-02"
	now, _ := time.Parse(shortForm, "2021-Feb-05")

	var deals []merkletree.Content

	for i := 1; i < 10; i++ {
		d := &Deal{
			DealID:      big.NewInt(now.Unix() + int64(i)),
			DataCID:     "datacid1234",
			PieceCID:    "piececid1234",
			Provider:    "fprovider1",
			StartEpoch:  big.NewInt(10),
			EndEpoch:    big.NewInt(2000),
			SignedEpoch: big.NewInt(50),
		}

		log.Debugw("generated deal", "deal", spew.Sdump(d))

		deals = append(deals, d)
	}

	tree, err := merkletree.NewTree(deals)
	if err != nil {
		panic(err)
	}

	mr := tree.MerkleRoot()
	MerkleRoot = common.BytesToHash(mr).Hex()
	log.Debugw("merkle root", "root", MerkleRoot)

	for i, d := range deals {
		a, _, err := tree.GetMerklePath(d)
		if err != nil {
			panic(err)
		}

		bytes := a[0]
		proof := common.BytesToHash(bytes).Hex()

		dd := d.(*Deal)
		dd.Proof = proof
		(deals[i].(*Deal)).Proof = proof
		log.Debug("got proof", "proof", dd.Proof)

		err = insertDeal(db.conn, dd)
		if err != nil {
			panic(err)
		}
	}

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
	var _dealID, _startEpoch, _endEpoch, _signedEpoch int64
	err := row.Scan(
		&_dealID,
		&deal.DataCID,
		&deal.PieceCID,
		&deal.Provider,
		&_startEpoch,
		&_endEpoch,
		&_signedEpoch,
		&deal.Proof)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	deal.DealID = big.NewInt(_dealID)
	deal.StartEpoch = big.NewInt(_startEpoch)
	deal.EndEpoch = big.NewInt(_endEpoch)
	deal.SignedEpoch = big.NewInt(_signedEpoch)
	return &deal, nil
}

func insertDeal(db *sql.DB, deal *Deal) error {
	insertSQL := `INSERT INTO ` +
		`deals(DealID, DataCID, PieceCID, Provider, StartEpoch, EndEpoch, SignedEpoch, Proof) VALUES ` +
		`(?, ?, ?, ?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(
		deal.DealID.Int64(),
		deal.DataCID,
		deal.PieceCID,
		deal.Provider,
		deal.StartEpoch.Int64(),
		deal.EndEpoch.Int64(),
		deal.SignedEpoch.Int64(),
		deal.Proof)
	return err
}
