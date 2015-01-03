package main

import (
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"labix.org/v2/mgo/bson"
)

// CollFile implements both Node and Handle for a collection.
type CollFile struct {
	Name string

	Dirent fuse.Dirent
	Fattr  fuse.Attr
}

func (f CollFile) Attr() fuse.Attr {
	slog.Noticef("CollFile.Attr() for: %+v", f)

	return fuse.Attr{Mode: os.ModeDir | 0400}
}

func (c CollFile) Lookup(fname string, intr fs.Intr) (fs.Node, fuse.Error) {
	slog.Noticef("CollFile[%s].Lookup(): %s\n", c.Name, fname)

	if !bson.IsObjectIdHex(fname) {
		return nil, fuse.ENOENT
	}

	db, s := getDb()
	defer s.Close()

	var f DocumentFile
	err := db.C(c.Name).FindId(bson.ObjectIdHex(fname)).One(&f)
	if err != nil {
		slog.Error(err)
		return nil, fuse.EIO
	}

	f.coll = c.Name

	return f, nil
}

func (c CollFile) ReadDir(intr fs.Intr) (ents []fuse.Dirent, ferr fuse.Error) {
	slog.Notice("CollFile.ReadDir():", c.Name)

	db, s := getDb()
	defer s.Close()

	iter := db.C(c.Name).Find(nil).Select(bson.M{"text": 0}).Iter()

	var f DocumentFile
	for iter.Next(&f) {
		ents = append(ents, fuse.Dirent{Name: f.Id.Hex(), Type: fuse.DT_File})
	}

	if err := iter.Err(); err != nil {
		slog.Error(err)
		return nil, fuse.EIO
	}

	return ents, nil
}
