package iorpc

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/minhajuddinkhan/iorung/cache/auth"
)

type JoinGameRequest struct {
	GameID string
	Token  string
}

//SetGameIDInToken sets game id against token
func (io *InterfaceRPC) SetGameIDInToken(req JoinGameRequest, out *bool) error {

	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(io.jwtSecret), nil
	})
	if err != nil {
		return fmt.Errorf("unable to decode jwt token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return fmt.Errorf("invalid token")
	}

	pl := auth.Player{
		PlayerID: claims["playerID"].(string),
		GameID:   req.GameID,
	}

	err = io.authRedis.Set(req.Token, pl)
	if err != nil {
		return err
	}
	*out = true

	return nil

}
