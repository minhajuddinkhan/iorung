package auth

import (
	"encoding/json"

	"github.com/gomodule/redigo/redis"
)

type Player struct {
	PlayerID string
	GameID   string
}

func (r *authRedis) Get(token string) (string, string, error) {

	conn, err := redis.DialURL(r.url)
	if err != nil {
		return "", "", err
	}
	defer conn.Close()

	s, err := redis.String(conn.Do("GET", token))
	if err != nil {
		return "", "", err
	}

	var pl Player
	err = json.Unmarshal([]byte(s), &pl)
	if err != nil {
		return "", "", err
	}

	return pl.GameID, pl.PlayerID, nil
}

//Set Set
func (r *authRedis) Set(token string, pl Player) error {
	conn, err := redis.DialURL(r.url)
	if err != nil {
		return err
	}
	b, err := json.Marshal(pl)
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Do("SET", token, b)
	return err
}
