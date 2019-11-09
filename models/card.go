package models

//Card playing card model
type Card struct {
	House  string `bson:"house"`
	Number int    `bson:"number"`
}
