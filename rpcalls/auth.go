package rpcalls

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/minhajuddinkhan/iorung"

	"github.com/minhajuddinkhan/iorung/cache/auth"
)

//AuthService AuthService
type AuthService struct {
	authRedis auth.Redis
}

//NewAuthListener creates a new auth listener
func NewAuthListener(conf *iorung.Conf) (net.Listener, error) {

	r, err := auth.NewAuthRedis(conf)
	if err != nil {
		return nil, err
	}
	service := new(AuthService)
	service.authRedis = r
	rpc.Register(service)
	rpc.HandleHTTP()
	fmt.Println("rpc calls listening on port: ", conf.IORungPort)
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.IORungPort))
	if err != nil {
		return nil, err
	}
	return l, nil
}

// Authenticate authenticates if a player has token in redis
func (as *AuthService) Authenticate(token string, out *bool) error {
	_, _, err := as.authRedis.Get(token)
	if err != nil {
		*out = true
		return nil
	}
	*out = false
	return nil
}

//Ping Ping
func (as *AuthService) Ping(in string, out *string) error {
	*out = in
	return nil
}
