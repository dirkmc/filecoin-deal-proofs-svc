package db

import (
	"database/sql"
)

const dsn = "./mock/sqlite-database.db"

type DB interface {
	DealByID(dealID uint64) (*Deal, error)
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
	DealID uint64
	DataCID string
	PieceCID string
	Provider string
	StartEpoch uint64
	EndEpoch uint64
	SignedEpoch uint64
}

func (db *liteDB) DealByID(dealID uint64) (*Deal, error) {
	//return &Deal{
	//	DealID:      dealID,
	//	DataCID:     "datacid1234",
	//	PieceCID:    "piececid1234",
	//	Provider:    "fprovider1",
	//	StartEpoch:  10,
	//	EndEpoch:    2000,
	//	SignedEpoch: 50,
	//}, nil

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