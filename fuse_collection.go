package main

import (
	"labix.org/v2/mgo/bson"
	"log"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type genericFile struct {
	Id bson.ObjectId `bson:"_id"`
}

// CollFile implements both Node and Handle for a collection.
type CollFile struct {
	Name string

	Dirent fuse.Dirent
	Fattr  fuse.Attr
}

func (f CollFile) Attr() fuse.Attr {
	log.Printf("CollFile.Attr() for: %+v", f)

	return collections[f.Name].Fattr
}

func (c CollFile) Lookup(fname string, intr fs.Intr) (fs.Node, fuse.Error) {
	log.Printf("CollFile[%s].Lookup(): %s\n", c.Name, fname)

	db, s := getDb()
	defer s.Close()

	var f interface{}
	err := db.C(c.Name).FindId(bson.ObjectIdHex(fname)).One(&f)
	if err != nil {
		logErr(err)
		return nil, fuse.EIO
	}

	log.Printf("Contents: %+v\n", f)

	return nil, fuse.ENOENT
}

func (c CollFile) ReadDir(intr fs.Intr) (ents []fuse.Dirent, ferr fuse.Error) {
	log.Println("CollFile.ReadDir():", c.Name)

	db, s := getDb()
	defer s.Close()

	iter := db.C(c.Name).Find(nil).Iter()

	var f genericFile
	for iter.Next(&f) {
		ents = append(ents, fuse.Dirent{Name: f.Id.Hex(), Type: fuse.DT_File})
	}

	if err := iter.Err(); err != nil {
		logErr(err)
		return nil, fuse.EIO
	}

	return ents, nil
}
