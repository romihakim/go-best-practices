package model

import "gopkg.in/mgo.v2/bson"

type Customer struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `json:"name" bson:"name"`
	Address string        `json:"address" bson:"address"`
	City    string        `json:"city" bson:"city"`
	State   string        `json:"state" bson:"state"`
	Zip     string        `json:"zip" bson:"zip"`
	Email   string        `json:"email" bson:"email"`
}
