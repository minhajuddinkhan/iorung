package auth

import (
	"github.com/gomodule/redigo/redis"
)

func (r *authRedis) Ping() error {

	conn, err := redis.DialURL(r.url)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}
