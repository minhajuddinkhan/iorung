package auth

import (
	"fmt"

	"github.com/minhajuddinkhan/iorung"
)

//Redis Redis
type Redis interface {
	Ping() error
	Get(token string) (gameID string, playerID string, err error)
	Set(token string, pl Player) error
}

type authRedis struct {
	url string
}

//NewAuthRedis returns connection of the authentication redis client
func NewAuthRedis(conf *iorung.Conf) (Redis, error) {
	if conf == nil {
		return nil, fmt.Errorf("redis err: nil configuration provided")
	}
	return &authRedis{url: conf.AuthRedis.RedisURL}, nil
}
