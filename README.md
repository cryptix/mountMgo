mountMgo
========

Reading a mongodb as a fusefs.


## working
* `ls mountpoint/` lists all collections in that database
* `ls mountpoint/collection` lists all documents with their `_id` as filename
* `cat mountpoint/collection/document` returns the document content with `json.MarhsalIndent`


## todo
- [ ] check some corner cases


## credits
* uses [bazil.org/fuse/](http://bazil.org/fuse/) for the fuse layer
* uses [labix.org/mgo](http://labix.org/mgo) for the mgo connection