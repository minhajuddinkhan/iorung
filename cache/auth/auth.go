package auth

import (
	"encoding/json"

	"github.com/gomodule/redigo/redis"
	"github.com/minhajuddinkhan/pattay"
)

type redisValue struct {
	gameID string
	player pattay.Player
}

func (r *authRedis) Get(token string) (string, pattay.Player, error) {

	conn, err := redis.DialURL(r.url)
	if err != nil {
		return "", nil, err
	}
	defer conn.Close()

	val, err := conn.Do("GET", token)
	if err != nil {
		return "", nil, err
	}

	b, err := redis.Bytes(val, err)
	if err != nil {
		return "", nil, err
	}

	var rd redisValue
	err = json.Unmarshal(b, &rd)
	if err != nil {
		return "", nil, err
	}

	return rd.gameID, rd.player, nil
}
