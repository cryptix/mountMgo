package main

import (
	"encoding/json"
	"time"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"labix.org/v2/mgo/bson"
)

// DocumentFile implements both Node and Handle for a document from a collection.
type DocumentFile struct {
	coll string
	Id   bson.ObjectId `bson:"_id"`

	Dirent fuse.Dirent
	Fattr  fuse.Attr
}

func (d DocumentFile) Attr() fuse.Attr {
	slog.Noticef("DocumentFile.Attr() for: %+v", d)

	now := time.Now()
	return fuse.Attr{
		Mode:  0400,
		Size:  42,
		Ctime: now,
		Atime: now,
		Mtime: now,
	}
}

func (d DocumentFile) Lookup(fname string, intr fs.Intr) (fs.Node, fuse.Error) {
	slog.Noticef("DocumentFile[%s].Lookup(): %s\n", d.coll, fname)

	return nil, fuse.ENOENT
}

func (d DocumentFile) ReadAll(intr fs.Intr) ([]byte, fuse.Error) {
	slog.Noticef("DocumentFile[%s].ReadAll(): %s\n", d.coll, d.Id)

	db, s := getDb()
	defer s.Close()

	var f interface{}
	err := db.C(d.coll).FindId(d.Id).One(&f)
	if err != nil {
		slog.Error(err)
		return nil, fuse.EIO
	}

	buf, err := json.MarshalIndent(f, "", "    ")
	if err != nil {
		slog.Error(err)
		return nil, fuse.EIO
	}

	return buf, nil
}
