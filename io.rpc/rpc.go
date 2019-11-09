package iorpc

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/minhajuddinkhan/iorung/cache/auth"
	"github.com/minhajuddinkhan/iorung/config"
)

//InterfaceRPC InterfaceRPC
type InterfaceRPC struct {
	authRedis auth.Redis
	jwtSecret string
}

//NewIOListener creates a new rpc listener for iorung
func NewIOListener(conf *config.Conf) (net.Listener, error) {

	r, err := auth.NewAuthRedis(conf)
	if err != nil {
		return nil, err
	}
	service := new(InterfaceRPC)
	service.authRedis = r
	service.jwtSecret = conf.JWTSecret
	rpc.Register(service)
	rpc.HandleHTTP()
	fmt.Println("rpc calls listening on port: ", conf.IORungPort)
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.IORungPort))
	if err != nil {
		return nil, err
	}
	return l, nil
}
