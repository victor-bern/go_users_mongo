package models

type Order struct {
	Data  string `bson:"data"`
	Items []Item `bson:"items"`
}
