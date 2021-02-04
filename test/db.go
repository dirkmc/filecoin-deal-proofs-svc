package test

import (
	"database/sql"
	"fmt"
	"github.com/dirkmc/filecoin-deal-proofs-svc/db"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
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
		DealID:      1,
		DataCID:     "datacid1",
		PieceCID:    "piececid1",
		Provider:    "provider1",
		StartEpoch:  10,
		EndEpoch:    1000,
		SignedEpoch: 50,
	})
	require.NoError(t, err)

	err = insertDeal(conn, &db.Deal{
		DealID:      2,
		DataCID:     "datacid2",
		PieceCID:    "piececid2",
		Provider:    "provider2",
		StartEpoch:  30,
		EndEpoch:    800,
		SignedEpoch: 70,
	})
	require.NoError(t, err)

	err = insertDeal(conn, &db.Deal{
		DealID:      3,
		DataCID:     "datacid3",
		PieceCID:    "piececid3",
		Provider:    "provider3",
		StartEpoch:  200,
		EndEpoch:    120,
		SignedEpoch: 250,
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
	insertSQL := `INSERT INTO `+
		`deals(DealID, DataCID, PieceCID, Provider, StartEpoch, EndEpoch, SignedEpoch) VALUES `+
		`(?, ?, ?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(
		deal.DealID,
		deal.DataCID,
		deal.PieceCID,
		deal.Provider,
		deal.StartEpoch,
		deal.EndEpoch,
		deal.SignedEpoch)
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
