package model

import "gopkg.in/mgo.v2/bson"

type Order struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	Date    string        `json:"date" bson:"date"`
	Name    string        `json:"name" bson:"name"`
	Product []string      `json:"product" bson:"product"`
	Total   float64       `json:"total" bson:"total"`
}
