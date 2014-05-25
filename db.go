package main

import (
	"fmt"
	"log"

	"labix.org/v2/mgo"
)

var getDb func() (*mgo.Database, *mgo.Session)

func initDb(dbHost, dbName string) {
	var err error

	url := fmt.Sprintf("mongodb://%s/%s", dbHost, dbName)
	mgoSession, err := mgo.Dial(url)
	checkFatal(err)

	log.Println("Dialed:", url)

	getDb = func() (*mgo.Database, *mgo.Session) {
		s := mgoSession.Clone()
		return s.DB(dbName), s
	}
}
