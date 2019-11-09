package models

import (
	"github.com/globalsign/mgo/bson"
	"github.com/minhajuddinkhan/pattay"
)

//Player player with cards
type Player struct {
	Cards    []pattay.Card `bson:"cards"`
	ID       bson.ObjectId `bson:"_id,omitempty"`
	PlayerID string        `bson:"player_id"`
	GameID   string        `bson:"game_id"`
}
