package iorung_test

import (
	"fmt"
	"net/rpc"
	"os"
	"testing"

	"github.com/minhajuddinkhan/iorung/cache/auth"
	"github.com/minhajuddinkhan/iorung/config"
	iorpc "github.com/minhajuddinkhan/iorung/io.rpc"
	"github.com/stretchr/testify/assert"
)

var host = "localhost"

// var port = os.Getenv("RPC_PORT")

func beforeEachRPC(t *testing.T) *rpc.Client {
	conf := config.New()

	conn, err := rpc.DialHTTP("tcp", fmt.Sprintf("%s:%d", host, conf.IORungPort))
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

var token string
var PlayerID = "1"

func TestIORungRPC_CanAddPlayer(t *testing.T) {

	req := iorpc.AddPlayerRequest{
		GameID:   "1",
		PlayerID: PlayerID,
	}
	conn := beforeEachRPC(t)
	err := conn.Call("InterfaceRPC.AddPlayer", req, &token)
	assert.Nil(t, err)

	url := os.Getenv("AUTH_REDIS_URL")
	r, err := auth.NewAuthRedis(&config.Conf{
		AuthRedis: config.Redis{
			RedisURL: url,
		},
	})
	playerID, gameID, err := r.Get(token)
	assert.Nil(t, err)
	assert.Equal(t, req.GameID, gameID)
	assert.Equal(t, req.PlayerID, playerID)

}

type Context struct {
	GameID   string `json:"game_id"`
	PlayerID string `json:"player_id"`
}

func TestIORungRPC_CanJoinGame(t *testing.T) {

	req := iorpc.AddPlayerRequest{
		GameID:   "55",
		PlayerID: "1",
	}
	conn := beforeEachRPC(t)
	err := conn.Call("InterfaceRPC.AddPlayer", req, &token)
	assert.Nil(t, err)

	joinReq := iorpc.JoinGameRequest{
		GameID: "55",
		Token:  token,
	}
	var done bool
	conn = beforeEachRPC(t)
	err = conn.Call("InterfaceRPC.SetGameIDInToken", joinReq, &done)
	assert.Nil(t, err)

	conf := config.New()
	r, err := auth.NewAuthRedis(&conf)

	gameID, playerID, err := r.Get(token)
	assert.Nil(t, err)
	assert.Equal(t, PlayerID, playerID)
	assert.Equal(t, gameID, req.GameID)

}
