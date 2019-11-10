package game

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
	"github.com/minhajuddinkhan/iorung/cache/auth"
	authManager "github.com/minhajuddinkhan/iorung/managers/auth"
)

//JoinGame JoinGame
func JoinGame(authRedis auth.Redis) func(s socketio.Conn, msg map[string]interface{}) {
	return func(s socketio.Conn, msg map[string]interface{}) {
		token := msg["token"].(string)

		mgr := authManager.New(authRedis)
		gameID, playerID, err := mgr.Authenticate(token)
		if err != nil {
			s.Emit("error", fmt.Errorf("invalid token: %s", token))
		}

		ctx := make(map[string]interface{})
		ctx["gameID"] = gameID
		ctx["playerID"] = playerID
		s.SetContext(ctx)

		s.Emit("player-joined", ctx)
	}

}
