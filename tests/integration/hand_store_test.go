package iorung_test

import (
	"testing"

	"github.com/minhajuddinkhan/iorung/config"
	"github.com/minhajuddinkhan/iorung/store/hand"
	"github.com/stretchr/testify/assert"
)

func TestHandStore_CanPing(t *testing.T) {
	conf := config.New()
	store := hand.NewHandStore(conf.DB.Connection)
	err := store.Ping()
	assert.Nil(t, err)
}
