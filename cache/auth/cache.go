package auth

import (
	"fmt"

	"github.com/minhajuddinkhan/iorung/config"
)

//Redis Redis
type Redis interface {
	Ping() error
	Get(token string) (gameID, playerID uint, err error)
	Set(token string, pl Player) error
	Delete(token string) error
}

type authRedis struct {
	url string
}

//NewAuthRedis returns connection of the authentication redis client
func NewAuthRedis(conf *config.Conf) (Redis, error) {
	if conf == nil {
		return nil, fmt.Errorf("redis err: nil configuration provided")
	}
	return &authRedis{url: conf.AuthRedis.RedisURL}, nil
}
