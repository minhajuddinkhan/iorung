package iorung_test

import (
	"testing"

	"github.com/minhajuddinkhan/iorung/cache/auth"
	"github.com/minhajuddinkhan/iorung/config"
	"github.com/stretchr/testify/assert"
)

func TestAuthRedis_PingShouldNotError(t *testing.T) {

	conf := config.New()
	r, err := auth.NewAuthRedis(&conf)
	assert.Nil(t, err)
	assert.NotNil(t, r)
	err = r.Ping()
	assert.Nil(t, err)

}
