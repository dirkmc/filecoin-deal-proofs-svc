package api

import (
	"github.com/dirkmc/filecoin-deal-proofs-svc/test"
	"testing"
)

func TestDB(t *testing.T) {
	test.Create(t, "../mock/sqlite-database.db")
}
