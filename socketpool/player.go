package socketpool

func (pp *playerpool) Join(id PlayerID, c PlayerConn) {
	pp.pool[id] = c
}

func (pp *playerpool) OnPlayer(id PlayerID) PlayerConn {
	return pp.pool[id]
}
