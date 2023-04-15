package main

import (
	"database/sql"
	"fmt"
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
	fmt.Println(os.Args)
	if err := app.AddBookmark(os.Args[2:]); err != nil {
		log.Fatal(err)
	}
}
