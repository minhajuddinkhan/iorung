package models

import (
	"github.com/globalsign/mgo/bson"
)

//Player player with cards
type Player struct {
	Cards    []Card        `bson:"cards"`
	ID       bson.ObjectId `bson:"_id,omitempty"`
	PlayerID string        `bson:"player_id"`
	GameID   string        `bson:"game_id"`
}
