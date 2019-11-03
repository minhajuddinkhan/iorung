package iorpc

// MAJOR REFACTORING REQUIRED. CANNOT PUSH CODE WITHOUT IT.
import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/minhajuddinkhan/iorung"

	"github.com/minhajuddinkhan/iorung/cache/auth"
)

//InterfaceRPC InterfaceRPC
type InterfaceRPC struct {
	authRedis auth.Redis
}

//NewIOListener creates a new rpc listener for iorung
func NewIOListener(conf *iorung.Conf) (net.Listener, error) {

	r, err := auth.NewAuthRedis(conf)
	if err != nil {
		return nil, err
	}
	service := new(InterfaceRPC)
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
func (as *InterfaceRPC) Authenticate(token string, out *bool) error {
	_, _, err := as.authRedis.Get(token)
	if err != nil {
		*out = true
		return nil
	}
	*out = false
	return nil
}

//Ping Ping
func (as *InterfaceRPC) Ping(in string, out *string) error {
	*out = in
	return nil
}
