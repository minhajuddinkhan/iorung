package iorung_test

import (
	"os"
	"testing"

	"github.com/minhajuddinkhan/iorung/store/hand"
	"github.com/stretchr/testify/assert"
)

func TestHandStore_CanPing(t *testing.T) {
	mongoConnStr := os.Getenv("MONGODB_URI")
	store := hand.NewHandStore(mongoConnStr)
	err := store.Ping()
	assert.Nil(t, err)
}
