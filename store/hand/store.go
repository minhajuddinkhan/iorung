package hand

type Store interface {
	Ping() error
}

type handStore struct {
	connectionString string
}

//NewHandStore returns a database client for hands collection
func NewHandStore(connStr string) Store {
	return &handStore{connectionString: connStr}
}
