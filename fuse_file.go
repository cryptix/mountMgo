package main

import (
	"log"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

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

func (c CollFile) ReadDir(intr fs.Intr) ([]fuse.Dirent, fuse.Error) {
	log.Println("CollFile.ReadDir():", c.Name)

	return nil, fuse.EIO

}
