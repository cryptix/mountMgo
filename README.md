mountMgo
========

Allows reading a mongodb as a fusefs.


## working
* `ls mountpoint/` lists all collections in that database
* `ls mountpoint/collection` lists all documents with their `_id` as filename
* `cat mountpoint/collection/document` returns the document content with `json.MarshalIndent`


## todo/ideas
- [x] collections are only read on startup
- [ ] check some corner cases
- [ ] documents in `xxx.index` collections don't have `_id` fields, so they aren't listed yet
- [ ] maybe experiment with mgo's GridFS(http://godoc.org/labix.org/v2/mgo#GridFS)

## credits
* uses [bazil.org/fuse/](http://bazil.org/fuse/) for the fuse layer
* uses [labix.org/mgo](http://labix.org/mgo) for the mgo connection
