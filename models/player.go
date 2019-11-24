package models

import (
	"github.com/globalsign/mgo/bson"
)

//Player player with cards
type Player struct {
	Cards    []Card        `json:"cards,omitempty" bson:"cards"`
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	PlayerID uint          `json:"player_id,omitempty" bson:"player_id"`
	GameID   uint          `json:"game_id" bson:"game_id"`
}
