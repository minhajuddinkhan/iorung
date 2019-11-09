package iorpc

//AddPlayerRequest AddPlayerRequest
type AddPlayerRequest struct {
	PlayerID string
	GameID   string
}

//AuthenticateResponse AuthenticateResponse
type AuthenticateResponse struct {
	GameID   string
	PlayerID string
}

//JoinGameRequest rpc request protocol for joining a game
type JoinGameRequest struct {
	GameID string
	Token  string
}

//DistributeCardsRequest rpc request protocol for distributing cards in a game
type DistributeCardsRequest struct {
	PlayerIds []string
	GameID    string
}

//DistributeCardsResponse DistributeCardsResponse
type DistributeCardsResponse struct {
	Players []Player
	GameID  string
}

type Player struct {
	Cards    []Card
	PlayerID string
	GameID   string
}
type Card struct {
	House  string
	Number int
}
