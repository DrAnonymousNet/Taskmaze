package storage

import (
	"fmt"
	"sync"

	"github.com/boltdb/bolt"
)

type TaskDB struct {
	DB *bolt.DB
	mu     sync.Mutex
}

var MyDB *TaskDB

func InitDB(dbName string) (*TaskDB, error) {
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Tasks"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error initializing db: %w", err)
	}

	MyDB = &TaskDB{DB: db}
	return MyDB, nil
}
