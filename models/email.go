package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"gopkg.in/mgo.v2"
)

type Email struct {
	ID			bson.ObjectId	`json:"id,omitempty" bson:"_id,omitempty"`
	Email		string			`json:"email" bson:"email"`
	Password	string			`json:"password" bson:"password"`
	Confirmed	bool			`json:"confirmed" bson:"confirmed"`
	CreatedAt	time.Time		`json:"created_at" bson:"created_at"`
	ConfirmedAt	time.Time		`json:"confirmed_at" bson:"confirmed_at"`
	ResetToken	string			`json:"reset_token" bson:"reset_token"`
	ResetAt		time.Time		`json:"reset_at" bson:"reset_at"`
}

// Index: (Email), unique
var EmailIndices = []mgo.Index{
	{
		Key: []string{"email"},
		Unique: true,
	},
}