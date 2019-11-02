package player

import (
	"github.com/globalsign/mgo"
)

func (hs *playerStore) Ping() error {

	db, err := mgo.Dial(hs.connectionString)
	if err != nil {
		return err
	}

	defer db.Close()
	return db.Ping()
}
