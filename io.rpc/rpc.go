package iorpc

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/minhajuddinkhan/iorung/cache/auth"
	"github.com/minhajuddinkhan/iorung/config"
	"github.com/minhajuddinkhan/iorung/socketpool"
	"github.com/minhajuddinkhan/iorung/store/player"
)

//InterfaceRPC InterfaceRPC
type InterfaceRPC struct {
	authRedis   auth.Redis
	playerStore player.Store
	jwtSecret   string
	gamepool    socketpool.GamePool
}

//NewIOListener creates a new rpc listener for iorung
func NewIOListener(r auth.Redis, playerStore player.Store, gamepool socketpool.GamePool, conf *config.Conf) (net.Listener, error) {

	service := new(InterfaceRPC)

	service.authRedis = r
	service.jwtSecret = conf.JWTSecret
	service.playerStore = playerStore
	service.gamepool = gamepool

	rpc.Register(service)
	rpc.HandleHTTP()
	fmt.Println("rpc calls listening on port: ", conf.IORungPort)
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.IORungPort))
	if err != nil {
		return nil, err
	}
	return l, nil
}
