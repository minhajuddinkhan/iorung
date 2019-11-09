package player

import (
	"github.com/globalsign/mgo"
	"github.com/minhajuddinkhan/iorung/models"
	"github.com/minhajuddinkhan/pattay"
)

func (s *playerStore) SetCardsAgainstPlayer(cards []pattay.Card, playerID, gameID string) error {

	session, err := mgo.Dial(s.connectionString)
	defer session.Clone()
	//TODO:: add custom errors for handling
	if err != nil {
		return err
	}

	dbCard := models.Player{Cards: cards, PlayerID: playerID, GameID: gameID}
	return session.DB(s.dbname).C(s.cardsCollection).Insert(&dbCard)

}
