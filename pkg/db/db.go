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

	dirPath := getMPMDir()

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
func getMPMDir() string {
	// 환경 변수 `MPM_DIR`이 설정되어 있으면 그걸 사용
	if customDir := os.Getenv("MPM_DIR"); customDir != "" {
		return customDir
	}

	// 기본적으로 홈 디렉토리에 저장
	home, err := os.UserHomeDir()
	if err != nil {
		panic("Failed to get home directory")
	}
	return filepath.Join(home, ".mpm")
}
