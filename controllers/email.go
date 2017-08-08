package controllers

import "gopkg.in/mgo.v2"

type EmailController struct {
	session *mgo.Session
}

func NewEmailController (sess *mgo.Session) *EmailController {
	return &EmailController{sess}
}
