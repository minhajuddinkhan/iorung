package socketpool

//PlayerPool manages socket connections of players
type PlayerPool interface {

	//Join joins a player's socket connection to the player pool
	Join(id PlayerID, connection PlayerConn)

	OnPlayer(id PlayerID) PlayerConn
}

//GamePool a pool of sockets
type GamePool interface {

	//JoinGame joins a player socket connection in the game pool
	JoinGame(gameID GameID, playerID PlayerID, connection PlayerConn) error
	//On returns the player pool of the game
	OnGame(id GameID) PlayerPool
}

//GamePool GamePool
type gamePool struct {
	pool map[GameID]PlayerPool
}

//NewGamePool NewGamePool
func NewGamePool() GamePool {

	return &gamePool{
		pool: make(map[GameID]PlayerPool),
	}
}

type playerpool struct {
	pool map[PlayerID]PlayerConn
}

//NewPlayerPool creates a new player pool
func NewPlayerPool() PlayerPool {
	return &playerpool{
		pool: make(map[PlayerID]PlayerConn),
	}
}
