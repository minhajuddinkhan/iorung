package iorpc

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/minhajuddinkhan/iorung/cache/auth"
	"github.com/minhajuddinkhan/rung"
)

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

//DistributeCards distributes cards amongst players
func (io *InterfaceRPC) DistributeCards(req DistributeCardsRequest, out *bool) error {

	game := rung.NewGame()
	if err := game.DistributeCards(); err != nil {
		return err
	}

	for i, player := range game.Players() {
		err := io.playerStore.SetCardsAgainstPlayer(
			player.CardsAtHand(),
			req.PlayerIds[i],
			req.GameID)
		if err != nil {
			return err
		}
	}

	*out = true
	return nil

}
