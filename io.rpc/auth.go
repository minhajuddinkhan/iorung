package iorpc

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/minhajuddinkhan/iorung/cache/auth"
)

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

//AddPlayer adds player to auth redis.
func (as *InterfaceRPC) AddPlayer(in AddPlayerRequest, out *string) error {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"playerID": in.PlayerID,
		"gameID":   in.GameID,
	})

	tokenString, err := token.SignedString([]byte(as.jwtSecret))
	if err != nil {
		return fmt.Errorf("cannot sign token. err: %v", err)
	}

	*out = tokenString
	player := auth.Player{
		GameID:   in.GameID,
		PlayerID: in.PlayerID,
	}

	err = as.authRedis.Set(tokenString, player)
	if err != nil {
		return fmt.Errorf("redis error: %v", err)
	}
	return nil
}
