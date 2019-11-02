package iorung_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

func TestSocket_CanConnect(t *testing.T) {

	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	assert.Nil(t, err)
	c, err := gosocketio.Dial(
		gosocketio.GetUrl(host, port, false),
		transport.GetDefaultWebsocketTransport(),
	)
	assert.Nil(t, err)
	assert.NotNil(t, c)
}
