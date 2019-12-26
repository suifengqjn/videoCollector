package store

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/storage"
)

var Store *DB

type DB struct {
	lvl  *leveldb.DB // Interface to the database itself
	Path string
}


// OpenDB opens a node database for storing and retrieving infos about known peers in the
// network. If no path is given an in-memory, temporary database is constructed.
func OpenDB(path string) (*DB, error) {
	if path == "" {
		return newMemoryDB()
	}
	return newPersistentDB(path)
}

// newMemoryNodeDB creates a new in-memory node database without a persistent backend.
func newMemoryDB() (*DB, error) {
	db, err := leveldb.Open(storage.NewMemStorage(), nil)
	if err != nil {
		return nil, err
	}
	return &DB{lvl: db}, nil
}

// newPersistentNodeDB creates/opens a leveldb backed persistent node database,
// also flushing its contents in case of a version mismatch.
func newPersistentDB(path string) (*DB, error) {
	opts := &opt.Options{OpenFilesCacheCapacity: 5}
	db, err := leveldb.OpenFile(path, opts)
	if _, iscorrupted := err.(*errors.ErrCorrupted); iscorrupted {
		db, err = leveldb.RecoverFile(path, nil)
	}
	if err != nil {
		return nil, err
	}

	return &DB{lvl: db}, nil
}

func (d *DB)Close()  {
	_ = d.lvl.Close()
}

func (d *DB) Save(key, value []byte) error {
	return d.lvl.Put(key, value, nil)
}

func (d *DB) Get(key []byte) []byte {
	ret, err := d.lvl.Has(key, nil)
	if !ret || err != nil {
		return nil
	}
	v, err :=  d.lvl.Get(key, nil)
	if err != nil {
		return nil
	}
	return v
}

func (d *DB)Has(key []byte) bool {
	ret, err := d.lvl.Has(key, nil)
	if err != nil {
		return false
	}
	return ret
}
