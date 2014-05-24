package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"bazil.org/fuse"
	"labix.org/v2/mgo"
)

var getDb func() (*mgo.Database, *mgo.Session)

var collections map[string]*CollFile

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

	names, err := mgoSession.DB(dbName).CollectionNames()
	checkFatal(err)

	collections = make(map[string]*CollFile, len(names))
	now := time.Now()
	for _, name := range names {
		collections[name] = &CollFile{
			Name: name,
			Fattr: fuse.Attr{
				Mode:  os.ModeDir | 0444,
				Ctime: now,
				Atime: now,
				Mtime: now,
			},
			Dirent: fuse.Dirent{Name: name, Type: fuse.DT_Dir},
		}
	}

}
