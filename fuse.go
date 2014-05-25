package main

import (
	"log"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

func mount(point string) {

	// startup mount
	c, err := fuse.Mount(point)
	checkFatal(err)
	defer c.Close()

	log.Println("Mounted: ", point)
	err = fs.Serve(c, mgoFS{})
	checkFatal(err)

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

	db, s := getDb()
	defer s.Close()

	names, err := db.CollectionNames()
	if err != nil {
		logErr(err)
		return nil, fuse.EIO
	}

	for _, collName := range names {
		if collName == name {
			return CollFile{Name: name}, nil
		}
	}

	return nil, fuse.ENOENT
}

func (d Dir) ReadDir(intr fs.Intr) ([]fuse.Dirent, fuse.Error) {
	log.Println("Dir.ReadDir():", d.name)

	db, s := getDb()
	defer s.Close()

	names, err := db.CollectionNames()
	if err != nil {
		logErr(err)
		return nil, fuse.EIO
	}

	ents := make([]fuse.Dirent, 0, len(names))

	for _, name := range names {
		ents = append(ents, fuse.Dirent{Name: name, Type: fuse.DT_Dir})
	}
	return ents, nil
}
