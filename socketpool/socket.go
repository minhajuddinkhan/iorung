package socketpool

import (
	"github.com/davecgh/go-spew/spew"
	socketio "github.com/googollee/go-socket.io"
	"github.com/minhajuddinkhan/iorung/models"
)

//PlayerConn the player socket connection
type PlayerConn interface {
	ReceiveInitialCards([]models.Card)
}

type socketconn struct {
	sio socketio.Conn
}

func NewPlayerConn(s socketio.Conn) PlayerConn {

	spew.Dump("NEW CONNECTION MADE!")
	return &socketconn{sio: s}
}

func (s *socketconn) ReceiveInitialCards(cards []models.Card) {
	s.sio.Emit("cards-init", cards)
}
