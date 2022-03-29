package main

import (
	"log"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/pgzisis/task/cmd"
	"github.com/pgzisis/task/db"
)

func main() {
	homeDir, err := homedir.Dir()
	if err != nil {
		log.Fatalln(err)
	}

	dbPath := filepath.Join(homeDir, "tasks.db")
	err = db.Init(dbPath)
	if err != nil {
		log.Fatalln(err)
	}

	cmd.RootCmd.Execute()
}
