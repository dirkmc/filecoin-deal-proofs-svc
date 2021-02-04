package test

import (
	"database/sql"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/dirkmc/filecoin-deal-proofs-svc/db"
	"github.com/stretchr/testify/require"
)

func Create(t *testing.T, dsn string) {
	file, err := os.Create(dsn)
	require.NoError(t, err)
	require.NoError(t, file.Close())

	conn, err := sql.Open("sqlite3", dsn)
	require.NoError(t, err)
	defer conn.Close()

	err = createTable(conn)
	require.NoError(t, err)

	err = insertDeal(conn, &db.Deal{
		DealID:      big.NewInt(1),
		DataCID:     "datacid1",
		PieceCID:    "piececid1",
		Provider:    "provider1",
		StartEpoch:  big.NewInt(10),
		EndEpoch:    big.NewInt(1000),
		SignedEpoch: big.NewInt(50),
	})
	require.NoError(t, err)

	err = insertDeal(conn, &db.Deal{
		DealID:      big.NewInt(2),
		DataCID:     "datacid2",
		PieceCID:    "piececid2",
		Provider:    "provider2",
		StartEpoch:  big.NewInt(30),
		EndEpoch:    big.NewInt(800),
		SignedEpoch: big.NewInt(70),
	})
	require.NoError(t, err)

	err = insertDeal(conn, &db.Deal{
		DealID:      big.NewInt(3),
		DataCID:     "datacid3",
		PieceCID:    "piececid3",
		Provider:    "provider3",
		StartEpoch:  big.NewInt(200),
		EndEpoch:    big.NewInt(120),
		SignedEpoch: big.NewInt(250),
	})
	require.NoError(t, err)

	deals, err := allDeals(conn)
	require.NoError(t, err)

	for _, d := range deals {
		fmt.Printf("id: %d, provider: %s\n", d.DealID, d.Provider)
	}
}

func createTable(db *sql.DB) error {
	createSQL := `CREATE TABLE deals (
		"DealID" integer,
		"DataCID" string,
		"PieceCID" string,
		"Provider" string,
		"StartEpoch" integer,
		"EndEpoch" integer,
		"SignedEpoch" integer
	  );`

	statement, err := db.Prepare(createSQL)
	if err != nil {
		return err
	}

	_, err = statement.Exec()
	return err
}

func insertDeal(db *sql.DB, deal *db.Deal) error {
	insertSQL := `INSERT INTO ` +
		`deals(DealID, DataCID, PieceCID, Provider, StartEpoch, EndEpoch, SignedEpoch) VALUES ` +
		`(?, ?, ?, ?, ?, ?, ?)`
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
		deal.SignedEpoch.Int64())
	return err
}

func allDeals(conn *sql.DB) ([]*db.Deal, error) {
	row, err := conn.Query("SELECT * FROM deals")
	if err != nil {
		return nil, err
	}

	var deals []*db.Deal
	for row.Next() {
		deal, err := db.RowToDeal(row)
		if err != nil {
			return nil, err
		}
		deals = append(deals, deal)
	}

	if row.Err() != nil {
		return nil, row.Err()
	}

	return deals, row.Close()
}
