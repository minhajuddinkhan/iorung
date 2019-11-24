package models

//Card playing card model
type Card struct {
	House  string `json:"house,omitempty" bson:"house"`
	Number int    `json:"number" bson:"number"`
}
