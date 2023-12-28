package main

import (
	"github.com/ywl0806/my-pj-manager/cmd"
	"github.com/ywl0806/my-pj-manager/pkg/db"
)

func main() {
	db.DbOnce.Do(db.InitalizeDB)
	cmd.Execute()

	defer db.CloseDB()
}
