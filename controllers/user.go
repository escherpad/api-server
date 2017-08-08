package controllers

import "gopkg.in/mgo.v2"

type UserController struct {
	session	*mgo.Session
}

func NewUserController (sess *mgo.Session) *UserController {
	return &UserController{sess}
}
