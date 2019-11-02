package player

//Store player store
type Store interface {
	Ping() error
}

type playerStore struct {
	connectionString string
}

//NewPlayerStore returns a database client for hands collection
func NewPlayerStore(connStr string) Store {
	return &playerStore{connectionString: connStr}
}
