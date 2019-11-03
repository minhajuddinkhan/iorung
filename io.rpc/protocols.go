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
