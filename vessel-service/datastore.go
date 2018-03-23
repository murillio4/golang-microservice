package main

import (
	mgo "gopkg.in/mgo.v2"
)

//CreateSession creates a mongo db session. connection to host
func CreateSession(host string) (*mgo.Session, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}
