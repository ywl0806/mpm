package db

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/boltdb/bolt"
)

var (
	DB     *bolt.DB
	DbOnce sync.Once
)
var models = []string{"PROJECTS"}

func InitalizeDB() {
	env := os.Getenv("ENV")

	var err error
	ex, err := os.Executable()

	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	dirPath := filepath.Join(exPath, ".mpm")
	if env == "dev" {
		dirPath = "tmp"
	}

	err = os.MkdirAll(dirPath, os.ModePerm)

	if err != nil {
		panic(err)
	}

	dbFilePath := filepath.Join(dirPath, "my.db")

	DB, err = bolt.Open(dbFilePath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Update(func(tx *bolt.Tx) error {
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
	return DB
}
func CloseDB() {
	DB.Close()
}
