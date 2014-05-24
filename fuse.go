package main

import (
	"log"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

var (
	newInode chan uint64
)

func mount(point string) {

	// init inode generator
	newInode = make(chan uint64)
	go func() {
		var counter uint64 = 2 // start with 2, 1 is dir root
		for {
			newInode <- counter
			counter += 1
		}
	}()

	// startup mount
	c, err := fuse.Mount(point)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	log.Println("Mounted: ", point)
	err = fs.Serve(c, mgoFS{})
	if err != nil {
		log.Fatal(err)
	}

	// check if the mount process has an error to report
	<-c.Ready
	checkFatal(c.MountError)
}

// mgoFS implements my mgo fuse filesystem
type mgoFS struct{}

func (mgoFS) Root() (fs.Node, fuse.Error) {
	log.Println("returning root node")
	return Dir{"Root"}, nil
}

// Dir implements both Node and Handle for the root directory.
type Dir struct {
	name string
}

func (d Dir) Attr() fuse.Attr {
	log.Println("Dir.Attr() for ", d.name)
	return fuse.Attr{Inode: 1, Mode: os.ModeDir | 0555}
}

func (Dir) Lookup(name string, intr fs.Intr) (fs.Node, fuse.Error) {
	log.Println("Dir.Lookup():", name)

	coll, ok := collections[name]
	if !ok {
		// might be..?
		return nil, fuse.ENOENT

	}

	return coll, nil
}

func (d Dir) ReadDir(intr fs.Intr) ([]fuse.Dirent, fuse.Error) {
	log.Println("Dir.ReadDir():", d.name)

	if d.name != "Root" {
		// list contents of collection
		return nil, fuse.ENOENT
	}

	// list collections
	db, s := getDb()
	defer s.Close()

	colls, err := db.CollectionNames()
	if err != nil {
		logErr(err)
		return nil, fuse.EIO
	}

	ents := make([]fuse.Dirent, 0, len(colls))

	for _, c := range colls {
		ents = append(ents, fuse.Dirent{Name: c, Type: fuse.DT_Dir})
	}
	return ents, nil
}
