package db

import (
	"fmt"
	"log"
	"sync"

	"github.com/boltdb/bolt"
)

var (
	Db     *bolt.DB
	DbOnce sync.Once
)
var models = []string{"PROJECTS"}

func InitalizeDB() {
	var err error
	Db, err = bolt.Open("tmp/my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		for _, model := range models {
			_, err = root.CreateBucketIfNotExists([]byte(model))
			if err != nil {
				return fmt.Errorf("could not create %s bucket: %v", model, err)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal("could not set up buckets")
	}

}

func GetDB() *bolt.DB {
	return Db
}
func CloseDB() {
	Db.Close()
}
