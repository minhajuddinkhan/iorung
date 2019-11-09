package player

import (
	"github.com/minhajuddinkhan/iorung/config"
	"github.com/minhajuddinkhan/pattay"
)

//Store player store
type Store interface {
	Ping() error
	SetCardsAgainstPlayer(cards []pattay.Card, playerID string, gameID string) error
}

type playerStore struct {
	connectionString string
	dbname           string
	cardsCollection  string
}

//NewPlayerStore returns a database client for hands collection
func NewPlayerStore(conf config.DB) Store {
	return &playerStore{
		connectionString: conf.Connection,
		dbname:           conf.DBName,
		cardsCollection:  "cards",
	}
}
