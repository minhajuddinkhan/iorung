package socketpool

func (gp *gamePool) JoinGame(gameID GameID, playerID PlayerID, c PlayerConn) error {

	if _, ok := gp.pool[gameID]; !ok {
		gp.pool[gameID] = NewPlayerPool()
	}
	gp.pool[gameID].Join(playerID, c)
	return nil
}

func (gp *gamePool) OnGame(id GameID) PlayerPool {
	return gp.pool[id]
}
