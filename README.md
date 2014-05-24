mountMgo
========

Reading a mongodb as a fusefs.


## working
* `ls mountpoint/` lists all collections in that database
* `ls mountpoint/collection` lists all documents with their `_id` as filename
* `cat mountpoint/collection/document` returns the document content with `json.MarhsalIndent`


## todo/ideas
- [ ] check some corner cases
- [ ] documents in `xxx.index` collections don't have `_id` fields, so they aren't listed yet
- [ ] maybe experiement with mgo's [GirdFs](http://godoc.org/labix.org/v2/mgo#GridFS)

## credits
* uses [bazil.org/fuse/](http://bazil.org/fuse/) for the fuse layer
* uses [labix.org/mgo](http://labix.org/mgo) for the mgo connection