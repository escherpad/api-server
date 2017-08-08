package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

type UserRole struct {
	Title		string			`json:"title" bson:"title"`
	BitMask		int				`json:"bit_mask" bson:"bit_mask"`
}

type UserSettings struct {
	CodeTheme		string			`json:"code_theme" bson:"code_theme"`
	CodeKeymap		string			`json:"code_keymap" bson:"code_keymap"`
	LatexCompiler	string			`json:"latex_compiler" bson:"latex_compiler"`
	LatexTimeout	string			`json:"latex_timeout" bson:"latex_timeout"`
	LatexPreamble	string			`json:"latex_preamble" bson:"latex_preamble"`
}
// User is model of users
type User struct {
	ID			bson.ObjectId	`json:"id,omitempty" bson:"_id,omitempty"`
	Username	string			`json:"username" bson:"username"`
	Team		bson.ObjectId	`json:"team" bson:"team"`
	Email		bson.ObjectId	`json:"email" bson:"email"`
	Role		UserRole		`json:"role" bson:"role"`
	Bindrs		[]bson.ObjectId	`json:"bindr" bson:"bindr"`
	Notes		[]bson.ObjectId	`json:"notes" bson:"notes"`
	Removed		bool			`json:"removed" bson:"removed"`
	Name		string			`json:"name" bson:"name"`
	WhatIDo		string			`json:"what_i_do" bson:"what_i_do"`
	cell		string			`json:"cell" bson:"cell"`
	Settings	UserSettings	`json:"settings" bson:"settings"`
}

// Index: (team, email), (team, username), unique
var UserIndices = []mgo.Index{
	{
		Key: []string{"team", "email"},
		Unique: true,
	},
	{
		Key: []string{"team", "username"},
		Unique: true,
	},
}