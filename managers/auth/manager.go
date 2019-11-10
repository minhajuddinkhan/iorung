package auth

import (
	auth "github.com/minhajuddinkhan/iorung/cache/auth"
)

//Manager authentication manager
type Manager interface {
	Authenticate(token string) (gameID string, playerID string, err error)
}

type manager struct {
	authRedis auth.Redis
}

func New(authRedis auth.Redis) Manager {
	return &manager{authRedis: authRedis}
}
