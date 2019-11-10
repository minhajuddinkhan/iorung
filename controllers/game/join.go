package game

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
	"github.com/minhajuddinkhan/iorung/cache/auth"
	authManager "github.com/minhajuddinkhan/iorung/managers/auth"
)

type Context struct {
	GameID   string `json:"game_id"`
	PlayerID string `json:"player_id"`
}

//JoinGame JoinGame
func JoinGame(authRedis auth.Redis) func(s socketio.Conn, msg map[string]interface{}) {
	return func(s socketio.Conn, msg map[string]interface{}) {
		token := msg["token"].(string)

		mgr := authManager.New(authRedis)
		gameID, playerID, err := mgr.Authenticate(token)
		if err != nil {
			s.Emit("error", fmt.Errorf("invalid token: %s", token))
		}

		ctx := Context{
			GameID:   gameID,
			PlayerID: playerID,
		}
		s.SetContext(ctx)
		s.Emit("pjoin", ctx)
	}

}
