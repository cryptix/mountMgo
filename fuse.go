package main

import (
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/cryptix/go/logging"
)

func mount(point string) {

	// startup mount
	c, err := fuse.Mount(point)
	logging.CheckFatal(err)
	defer c.Close()

	slog.Noticef("Mounted: ", point)
	err = fs.Serve(c, mgoFS{})
	logging.CheckFatal(err)

	// check if the mount process has an error to report
	<-c.Ready
	logging.CheckFatal(c.MountError)
}

// mgoFS implements my mgo fuse filesystem
type mgoFS struct{}

func (mgoFS) Root() (fs.Node, fuse.Error) {
	slog.Noticef("returning root node")
	return Dir{"Root"}, nil
}

// Dir implements both Node and Handle for the root directory.
type Dir struct {
	name string
}

func (d Dir) Attr() fuse.Attr {
	slog.Noticef("Dir.Attr() for ", d.name)
	return fuse.Attr{Inode: 1, Mode: os.ModeDir | 0555}
}

func (Dir) Lookup(name string, intr fs.Intr) (fs.Node, fuse.Error) {
	slog.Noticef("Dir.Lookup():", name)

	db, s := getDb()
	defer s.Close()

	names, err := db.CollectionNames()
	if err != nil {
		slog.Error(err)
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
	slog.Noticef("Dir.ReadDir():", d.name)

	db, s := getDb()
	defer s.Close()

	names, err := db.CollectionNames()
	if err != nil {
		slog.Error(err)
		return nil, fuse.EIO
	}

	ents := make([]fuse.Dirent, 0, len(names))

	for _, name := range names {
		ents = append(ents, fuse.Dirent{Name: name, Type: fuse.DT_Dir})
	}
	return ents, nil
}
