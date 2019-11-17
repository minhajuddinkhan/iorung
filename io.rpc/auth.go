package iorpc

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/minhajuddinkhan/iorung/cache/auth"
	authManager "github.com/minhajuddinkhan/iorung/managers/auth"
)

// Authenticate authenticates if a player has token in redis
func (as *InterfaceRPC) Authenticate(token string, out *AuthenticateResponse) error {
	mgr := authManager.New(as.authRedis)

	gameID, playerID, err := mgr.Authenticate(token)
	if err != nil {
		return fmt.Errorf("unable to authenticate. err:%v", err)
	}

	*out = AuthenticateResponse{
		GameID:   gameID,
		PlayerID: playerID,
	}
	return nil
}

//AddPlayer adds player to auth redis.
func (as *InterfaceRPC) AddPlayer(in AddPlayerRequest, out *string) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"playerID": in.PlayerID,
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"iat":      time.Now().Unix(),
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

//Logout Logout
func (as *InterfaceRPC) Logout(req LogoutRequest, out *bool) error {
	return as.authRedis.Delete(req.Token)
}
