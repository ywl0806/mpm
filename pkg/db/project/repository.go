package project

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/ywl0806/my-pj-manager/pkg/db"
)

func Bucket(tx *bolt.Tx) *bolt.Bucket {
	return tx.Bucket([]byte("DB")).Bucket([]byte("PROJECTS"))
}
func List() (projects []Project, err error) {

	err = db.Db.View(func(tx *bolt.Tx) error {
		bucket := Bucket(tx)
		cursor := bucket.Cursor()
		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			var pj Project
			db.Deserialize(value, &pj)
			projects = append(projects, pj)
		}
		return nil
	})
	return projects, err
}

func Add(newPj Project) error {

	err := db.Db.Update(func(tx *bolt.Tx) error {
		var err error
		if err != nil {
			return fmt.Errorf("transaction start error: %v", err)
		}

		b := Bucket(tx)

		if IsProjectNameDuplicate(newPj.Name, b) {
			return fmt.Errorf("duplicated name: %s", newPj.Name)
		}

		err = b.Put([]byte(newPj.Name), db.Serialize(newPj))

		if err != nil {
			return fmt.Errorf("could not insert project: %v", err)
		}
		return err
	})

	return err
}

// Args{b} is optional and can be set to nil
func IsProjectNameDuplicate(name string, b *bolt.Bucket) bool {

	if b == nil {

		tx, err := db.Db.Begin(true)
		if err != nil {
			log.Fatalln("transaction start error")
		}

		b = Bucket(tx)
	}

	return b.Get([]byte(name)) != nil
}
