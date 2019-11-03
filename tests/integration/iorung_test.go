package iorung_test

import (
	"fmt"
	"net/rpc"
	"os"
	"testing"

	"github.com/minhajuddinkhan/iorung"
	"github.com/minhajuddinkhan/iorung/cache/auth"
	iorpc "github.com/minhajuddinkhan/iorung/io.rpc"
	"github.com/stretchr/testify/assert"
)

var host = "localhost"
var port = os.Getenv("RPC_PORT")

func beforeEachRPC(t *testing.T) *rpc.Client {
	conn, err := rpc.DialHTTP("tcp", fmt.Sprintf("%s:%s", host, port))
	assert.Nil(t, err)
	return conn

}
func TestIORungRPC_CanPing(t *testing.T) {

	conn := beforeEachRPC(t)

	input := "hello"
	var out string
	err := conn.Call("InterfaceRPC.Ping", input, &out)
	assert.Nil(t, err)
	assert.NotNil(t, conn)
	assert.Equal(t, input, out)
}

func TestIORungRPC_CanAddPlayer(t *testing.T) {

	req := iorpc.AddPlayerRequest{
		GameID:   "1",
		PlayerID: "1",
	}
	conn := beforeEachRPC(t)
	var token string
	err := conn.Call("InterfaceRPC.AddPlayer", req, &token)
	assert.Nil(t, err)

	url := os.Getenv("AUTH_REDIS_URL")
	r, err := auth.NewAuthRedis(&iorung.Conf{
		AuthRedis: iorung.Redis{
			RedisURL: url,
		},
	})

	playerID, gameID, err := r.Get(token)
	assert.Nil(t, err)
	assert.Equal(t, req.GameID, gameID)
	assert.Equal(t, req.PlayerID, playerID)

}
