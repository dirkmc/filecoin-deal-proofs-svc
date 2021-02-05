package db

import (
	"math/big"
	"testing"

	"github.com/cbergoon/merkletree"
	"github.com/ethereum/go-ethereum/common"
)

func TestCalculateHash(t *testing.T) {
	d1 := Deal{
		DataCID:     "datacid1234",
		PieceCID:    "piececid1234",
		DealID:      big.NewInt(150505),
		Provider:    "fprovider1",
		StartEpoch:  big.NewInt(10),
		EndEpoch:    big.NewInt(2000),
		SignedEpoch: big.NewInt(50),
	}

	d2 := Deal{
		DataCID:     "datacid2345",
		PieceCID:    "piececid2345",
		DealID:      big.NewInt(267775),
		Provider:    "fprovider2",
		StartEpoch:  big.NewInt(100),
		EndEpoch:    big.NewInt(4000),
		SignedEpoch: big.NewInt(250),
	}

	var deals []merkletree.Content

	deals = append(deals, d1)
	deals = append(deals, d2)

	tree, err := merkletree.NewTree(deals)
	if err != nil {
		panic(err)
	}

	a, _, err := tree.GetMerklePath(d1)
	if err != nil {
		panic(err)
	}

	bytes := a[0]
	got := common.BytesToHash(bytes).Hex()

	expected := "0xe476bc2413fc2aff9f04a6ad695bc158b5420a6d98f6691cfb7ce9a9ec2cdcf9"

	if got != expected {
		t.Fatalf("expected to get hash %s, but got %s", expected, got)
	}
}
