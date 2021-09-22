package models

import (
	"time"
)

type Order struct {
	Data  time.Time `bson:"data"`
	Items []Item    `bson:"items"`
}
