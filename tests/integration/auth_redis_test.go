package iorung_test

import (
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/minhajuddinkhan/iorung"
	"github.com/minhajuddinkhan/iorung/cache/auth"
	"github.com/stretchr/testify/assert"
)

func TestAuthRedis_PingShouldNotError(t *testing.T) {

	url := os.Getenv("AUTH_REDIS_URL")
	r, err := auth.NewAuthRedis(&iorung.Conf{
		AuthRedis: iorung.Redis{
			RedisURL: url,
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, r)
	err = r.Ping()
	spew.Dump(err)
	assert.Nil(t, err)
}
