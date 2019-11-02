package iorung_test

import (
	"os"
	"testing"

	"github.com/minhajuddinkhan/iorung/store/player"
	"github.com/stretchr/testify/assert"
)

func TestPlayerStore_CanPing(t *testing.T) {
	mongoConnStr := os.Getenv("MONGODB_URI")
	store := player.NewPlayerStore(mongoConnStr)
	err := store.Ping()
	assert.Nil(t, err)
}
