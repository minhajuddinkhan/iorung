package iorung_test

import (
	"os"
	"testing"

	"github.com/minhajuddinkhan/iorung/cache/socket"
	"github.com/minhajuddinkhan/iorung/config"
	"github.com/stretchr/testify/assert"
)

func TestSocketRedis_PingShouldNotError(t *testing.T) {

	url := os.Getenv("SOCKET_REDIS_URL")

	r, err := socket.NewSocketRedis(&config.Conf{
		SocketRedis: config.Redis{
			RedisURL: url,
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, r)
	err = r.Ping()
	assert.Nil(t, err)
}
