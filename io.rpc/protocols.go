package iorpc

//AddPlayerRequest AddPlayerRequest
type AddPlayerRequest struct {
	PlayerID uint
	GameID   uint
}

//AuthenticateResponse AuthenticateResponse
type AuthenticateResponse struct {
	GameID   uint
	PlayerID uint
}

//JoinGameRequest rpc request protocol for joining a game
type JoinGameRequest struct {
	GameID uint
	Token  string
}

//DistributeCardsRequest rpc request protocol for distributing cards in a game
type DistributeCardsRequest struct {
	PlayerIds []uint
	GameID    uint
}

type Player struct {
	Cards    []Card
	PlayerID uint
	GameID   uint
}
type Card struct {
	House  string
	Number int
}
