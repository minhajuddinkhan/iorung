package game

import (
	"github.com/davecgh/go-spew/spew"
	socketio "github.com/googollee/go-socket.io"
	"github.com/minhajuddinkhan/iorung/socketpool"
)

type Context struct {
	GameID   uint `json:"game_id"`
	PlayerID uint `json:"player_id"`
}

//JoinGame JoinGame
func (ctrl *Controller) JoinGame(s socketio.Conn, msg map[string]interface{}) {

	player, ok := s.Context().(Ctx)
	if !ok {
		s.Emit("join:error", "player not found in context")
		return
	}
	if _, ok := msg["game_id"]; !ok {
		s.Emit("join:error", "game_id not found")
		return
	}
	spew.Dump(msg["game_id"])
	gameID, ok := msg["game_id"].(float64)
	if !ok {
		s.Emit("join:error", "game_id is not an unsigned integer")
		return
	}
	gid := socketpool.GameID(gameID)
	pid := socketpool.PlayerID(player.PlayerID)
	err := ctrl.gamepool.JoinGame(gid, pid, socketpool.NewPlayerConn(s))

	if err != nil {
		s.Emit("join:error", err.Error())
		return
	}
	s.Emit("join:done", map[string]interface{}{"done": true})
}
