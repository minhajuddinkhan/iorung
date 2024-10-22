package iorung_test

import (
	"os"
	"strconv"
	"testing"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"github.com/stretchr/testify/assert"
)

func TestSocket_CanConnect(t *testing.T) {

	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	assert.Nil(t, err)
	c, err := gosocketio.Dial(
		gosocketio.GetUrl(host, port, false),
		transport.GetDefaultWebsocketTransport(),
	)
	defer c.Close()
	assert.Nil(t, err)
	assert.NotNil(t, c)
}
