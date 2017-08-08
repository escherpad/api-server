package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

type Team struct {
	ID		bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name	string	`json:"name" bson:"name"`
	Domain	string	`json:"domain" bson:"domain"`
}

// Index (domain), unique
var TeamIndices = []mgo.Index{
	{
		Key: []string{"domain"},
		Unique: true,
	},
}