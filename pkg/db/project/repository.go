package project

import (
	"errors"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/ywl0806/mpm/pkg/db"
)

func Bucket(tx *bolt.Tx) *bolt.Bucket {
	return tx.Bucket([]byte("DB")).Bucket([]byte("PROJECTS"))
}
func FindByName(name string) (pj Project, err error) {
	err = db.DB.View(func(tx *bolt.Tx) (viewErr error) {
		bucket := Bucket(tx)
		rawPj := bucket.Get([]byte(name))

		if rawPj == nil {
			return errors.New(" Not Found")
		}

		viewErr = db.Deserialize(rawPj, &pj)
		return viewErr
	})

	return pj, err
}
func List() (projects []Project, err error) {

	err = db.DB.View(func(tx *bolt.Tx) error {
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

	err := db.DB.Update(func(tx *bolt.Tx) error {
		var err error

		b := Bucket(tx)

		if CheckIsProjectExist(newPj.Name, b) {
			return fmt.Errorf("duplicated name: %s", newPj.Name)
		}

		err = b.Put([]byte(newPj.Name), db.Serialize(&newPj))

		if err != nil {
			return fmt.Errorf("could not insert project: %v", err)
		}
		return err
	})

	return err
}

func Update(name string, update *UpdateProject) (err error) {

	err = db.DB.Update(func(tx *bolt.Tx) (updateErr error) {
		bucket := Bucket(tx)

		var pj Project

		rawPj := bucket.Get([]byte(name))

		if rawPj == nil {
			return errors.New(" Not Found")
		}

		updateErr = db.Deserialize(rawPj, &pj)

		if updateErr != nil {
			return updateErr
		}

		if update.Path != nil {
			pj.Path = *update.Path
		}
		if update.Directories != nil {
			pj.Directories = *update.Directories
		}
		if update.Last_use_at != nil {
			pj.Last_use_at = *update.Last_use_at
		}
		if update.Usage != nil {
			pj.Usage = *update.Usage
		}

		updateErr = bucket.Put([]byte(name), db.Serialize(&pj))

		return updateErr
	})
	return err
}

func Delete(name string) error {
	err := db.DB.Update(func(tx *bolt.Tx) error {
		bucket := Bucket(tx)

		err := bucket.Delete([]byte(name))

		return err
	})

	return err
}

// Args{b} is optional and can be set to nil
func CheckIsProjectExist(name string, b *bolt.Bucket) bool {

	if b == nil {

		tx, err := db.DB.Begin(true)
		if err != nil {
			log.Fatalln("transaction start error")
		}

		b = Bucket(tx)
	}

	return b.Get([]byte(name)) != nil
}
