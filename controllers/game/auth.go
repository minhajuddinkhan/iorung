package game

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
	"github.com/minhajuddinkhan/iorung/errs"
	authManager "github.com/minhajuddinkhan/iorung/managers/auth"
	"github.com/minhajuddinkhan/iorung/socketpool"
)

var AuthErrEvent string = "authenticate:error"
var AuthSuccessEvent string = "authenticate:done"

func (ctrl *Controller) Authenticate(s socketio.Conn, msg map[string]interface{}) {

	token := msg["token"].(string)
	mgr := authManager.New(ctrl.authRedis)
	gameID, playerID, err := mgr.Authenticate(token)
	if err != nil {
		s.Emit(AuthErrEvent, fmt.Errorf("invalid token: %s", token))
	}

	if gameID == 0 {
		s.Emit(AuthErrEvent, fmt.Errorf("empty game ID"))
	}
	if playerID == 0 {
		s.Emit(AuthErrEvent, fmt.Errorf("empty player ID"))
	}

	ctx := Ctx{PlayerID: playerID, GameID: gameID}
	s.SetContext(ctx)

	player, err := ctrl.playerStore.GetPlayer(playerID)

	if err != nil {
		switch err.(type) {
		case *errs.ErrPlayerNotFound:
			s.Emit(AuthSuccessEvent, player)
			return
		default:

			s.Emit(AuthErrEvent, err.Error())
			return
		}

	}
	socketpool.NewPlayerConn(s).ReceiveInitialCards(player.Cards)

	s.Emit(AuthSuccessEvent, player)
	return

}
