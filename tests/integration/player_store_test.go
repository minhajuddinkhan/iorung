package iorung_test

import (
	"testing"

	"github.com/minhajuddinkhan/iorung/config"
	"github.com/minhajuddinkhan/iorung/store/player"
	"github.com/stretchr/testify/assert"
)

func TestPlayerStore_CanPing(t *testing.T) {

	conf := config.New()
	store := player.NewPlayerStore(conf.DB)
	err := store.Ping()
	assert.Nil(t, err)
}
