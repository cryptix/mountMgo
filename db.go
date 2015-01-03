package main

import (
	"fmt"

	"labix.org/v2/mgo"
)

var getDb func() (*mgo.Database, *mgo.Session)

func initDb(dbHost, dbName string) {
	url := fmt.Sprintf("mongodb://%s/%s", dbHost, dbName)
	mgoSession, err := mgo.Dial(url)
	if err != nil {
		slog.Fatal(err)
	}

	slog.Info("Dialed:", url)

	getDb = func() (*mgo.Database, *mgo.Session) {
		s := mgoSession.Clone()
		return s.DB(dbName), s
	}
}
