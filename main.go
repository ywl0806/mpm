package main

import (
	"github.com/ywl0806/mpm/cmd"
	"github.com/ywl0806/mpm/pkg/db"
)

func main() {
	db.DbOnce.Do(db.InitalizeDB)
	cmd.Execute()

	defer db.CloseDB()
}
