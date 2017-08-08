package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"gopkg.in/mgo.v2"
)

type RefreshToken struct {
	ID			bson.ObjectId	`json:"id,omitempty" bson:"_id,omitempty"`
	Token		string			`json:"token" bson:"token"`
	User		bson.ObjectId	`json:"user" bson:"user"`
	ExpireAt	time.Time		`json:"expire_at" bson:"expire_at"`
	DeviceID	string			`json:"device_id" bson:"device_id"`
	CreatedAt	time.Time		`json:"created_at" bson:"created_at"`
	Expires		string			`json:"expires" bson:"expires"`
	UsedAt		time.Time		`json:"used_at" bson:"used_at"`
}

// Index (device_id, user), sparse unique
var RefreshTokenIndices = []mgo.Index{
	{
		Key: []string{"device_id", "user"},
		Unique: true,
		Sparse: true,
	},
}