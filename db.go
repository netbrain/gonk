package main

import "labix.org/v2/mgo"

var session *mgo.Session

func initDB() {
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
}
