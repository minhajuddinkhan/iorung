package game

import (
	"github.com/minhajuddinkhan/iorung/cache/auth"
	"github.com/minhajuddinkhan/iorung/socketpool"
	"github.com/minhajuddinkhan/iorung/store/player"
)

type Controller struct {
	authRedis   auth.Redis
	playerStore player.Store
	gamepool    socketpool.GamePool
}

//NewGameCtrl NewGameCtrl
func NewGameCtrl(r auth.Redis, ps player.Store, gp socketpool.GamePool) *Controller {
	return &Controller{
		authRedis:   r,
		playerStore: ps,
		gamepool:    gp,
	}
}
