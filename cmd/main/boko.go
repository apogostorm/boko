package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/apogostorm/boko/pkg/app"
	"github.com/apogostorm/boko/pkg/bookmarks"
	_ "github.com/mattn/go-sqlite3"
)

func getApp() *app.App {
	dirname, err := os.UserHomeDir()
	dbFileName := dirname + "/.bookmarks/bookmarks.db"
	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	r := bookmarks.NewRepo(db)
	return &app.App{
		BookmarkRepo: r,
	}
}

func main() {
	app := getApp()
	if err := app.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
