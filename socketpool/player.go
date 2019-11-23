package socketpool

import (
	"fmt"
)

func (pp *playerpool) Join(id PlayerID, c PlayerConn) {
	pp.pool[id] = c
	// pp.pool[id].Emit("done", map[string]interface{}{"joined": true})
	fmt.Println("players joined", len(pp.pool))
}

func (pp *playerpool) ReceiveInitialCards(id PlayerID, cards []interface{}) {

}

func (pp *playerpool) OnPlayer(id PlayerID) PlayerConn {
	return pp.pool[id]
}
