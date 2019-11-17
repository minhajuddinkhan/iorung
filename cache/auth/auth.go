package auth

import (
	"encoding/json"

	"github.com/gomodule/redigo/redis"
)

type Player struct {
	PlayerID uint
	GameID   uint
}

func (r *authRedis) Get(token string) (uint, uint, error) {

	conn, err := redis.DialURL(r.url)
	if err != nil {
		return 0, 0, err
	}
	defer conn.Close()

	s, err := redis.String(conn.Do("GET", token))
	if err != nil {
		return 0, 0, err
	}

	var pl Player
	err = json.Unmarshal([]byte(s), &pl)
	if err != nil {
		return 0, 0, err
	}

	return pl.GameID, pl.PlayerID, nil
}

//Set Set
func (r *authRedis) Set(token string, pl Player) error {
	conn, err := redis.DialURL(r.url)
	if err != nil {
		return err
	}
	defer conn.Close()
	b, err := json.Marshal(pl)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", token, b)
	return err
}

func (r *authRedis) Delete(token string) error {

	conn, err := redis.DialURL(r.url)
	if err != nil {
		return err
	}

	defer conn.Close()
	if _, err := conn.Do("DEL", token); err != nil {
		return err
	}

	return nil

}
