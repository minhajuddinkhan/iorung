package models

import (
	"github.com/globalsign/mgo/bson"
)

//Player player with cards
type Player struct {
	Cards    []Card        `bson:"cards"`
	ID       bson.ObjectId `bson:"_id,omitempty"`
	PlayerID uint          `bson:"player_id"`
	GameID   uint          `bson:"game_id"`
}
