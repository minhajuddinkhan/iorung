package iorung_test

import (
	"fmt"
	"net/rpc"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIORungRPC_CanPing(t *testing.T) {

	host := "localhost"
	port := os.Getenv("RPC_PORT")
	input := "hello"
	var out string
	conn, err := rpc.DialHTTP("tcp", fmt.Sprintf("%s:%s", host, port))
	assert.Nil(t, err)

	err = conn.Call("AuthService.Ping", input, &out)

	assert.Nil(t, err)
	assert.NotNil(t, conn)
	assert.Equal(t, input, out)
}