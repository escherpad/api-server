package services

import (
	"gopkg.in/mgo.v2"
	m "github.com/escherpad/api-server/models"
)

type DS struct {
	Session *mgo.Session
}


func (ds *DS)Init(addr string, db string) error{
	session, err := mgo.DialWithInfo(&mgo.DialInfo{Addrs: []string{addr}, Database: db})
	if err != nil {
		return err
	}
	collection := []string{"email", "refresh_token", "team", "user", "comment"}
	d := session.DB("")
	index := [][]mgo.Index{m.EmailIndices, m.RefreshTokenIndices, m.TeamIndices, m.UserIndices, m.CommentIndices}
	for i, _ := range collection {
		for _, ind := range index[i] {
			d.C(collection[i]).EnsureIndex(ind)
		}
	}
	ds.Session = session
	return err
}

func (ds *DS)Close() {
	ds.Session.Close()
}

// Not sure if it is a good design
var DataStore DS = DS{}