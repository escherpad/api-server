package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"time"
)

type Comment struct {
	ID				bson.ObjectId	`json:"id,omitempty" bson:"_id,omitempty"`
	Selection		interface{}		`json:"selection" bson:"selection"`
	Text			string			`json:"text" bson:"text"`
	Users			bson.ObjectId	`json:"users" bson:"users"`
	Authors			bson.ObjectId	`json:"authors" bson:"authors"`
	TimeCreated		time.Time		`json:"time_created" bson:"time_created"`
	TimeModified	time.Time		`json:"time_modified" bson:"time_modified"`
	TimeUpdated		time.Time		`json:"time_updated" bson:"time_updated"`
	Archived		bool			`json:"archived" bson:"archived"`
}

// Index (Users.Username, -TimeModified), (Authors.Username, -TimeModified)

var CommentIndices = []mgo.Index{
	{
		Key: []string{"Users", "-TimeModified)"},
	},
	{
		Key: []string{"Authors", "-TimeModified)"},
	},
}