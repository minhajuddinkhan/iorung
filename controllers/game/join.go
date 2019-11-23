package game

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
	"github.com/minhajuddinkhan/iorung/cache/auth"
	authManager "github.com/minhajuddinkhan/iorung/managers/auth"
	"github.com/minhajuddinkhan/iorung/socketpool"
)

type Context struct {
	GameID   uint `json:"game_id"`
	PlayerID uint `json:"player_id"`
}

//JoinGame JoinGame
func JoinGame(authRedis auth.Redis, gamepool socketpool.GamePool) func(s socketio.Conn, msg map[string]interface{}) {
	return func(s socketio.Conn, msg map[string]interface{}) {
		token := msg["token"].(string)
		fmt.Println("RECEIVED!", token)
		mgr := authManager.New(authRedis)
		gameID, playerID, err := mgr.Authenticate(token)
		if err != nil {
			s.Emit("error", fmt.Errorf("invalid token: %s", token))
		}

		if gameID == 0 {
			s.Emit("error", fmt.Errorf("empty game ID"))
		}
		if playerID == 0 {
			s.Emit("erro", fmt.Errorf("empty player ID"))
		}
		ctx := Context{
			GameID:   gameID,
			PlayerID: playerID,
		}
		s.SetContext(ctx)
		gamepool.JoinGame(socketpool.GameID(gameID), socketpool.PlayerID(playerID), s)
	}

}
