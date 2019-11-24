package player

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/minhajuddinkhan/iorung/errs"
	"github.com/minhajuddinkhan/iorung/models"
)

func (s *playerStore) SetCardsAgainstPlayer(cards []models.Card, playerID, gameID uint) error {

	session, err := mgo.Dial(s.connectionString)
	defer session.Clone()
	//TODO:: add custom errors for handling
	if err != nil {
		return err
	}

	dbCard := models.Player{Cards: cards, PlayerID: playerID, GameID: gameID}
	return session.DB(s.dbname).C(s.cardsCollection).Insert(&dbCard)

}

func (s *playerStore) GetPlayer(playerID uint) (*models.Player, error) {

	session, err := mgo.Dial(s.connectionString)
	defer session.Clone()
	//TODO:: add custom errors for handling
	if err != nil {
		return nil, err
	}
	var player models.Player
	err = session.DB(s.dbname).C(s.cardsCollection).Find(bson.M{"player_id": playerID}).One(&player)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, &errs.ErrPlayerNotFound{ID: playerID}
		}
		return nil, err
	}

	return &player, err

}
