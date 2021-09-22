package models

type Item struct {
	Name     string  `bson:"name"`
	Quantity int     `bson:"qtd"`
	Price    float64 `bson:"price"`
}
