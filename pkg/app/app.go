package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/apogostorm/boko/pkg/bookmarks"
)

type App struct {
	BookmarkRepo bookmarks.Repo
}

const (
	addHelpMessage     = "Usage: boko add <url> <name> [tag1 tag2 ... tagn]"
	findHelpMessage    = "Usage: boko find <name-or-tag> | (--name|-n <name>, --tag|-t <tag>)"
	generalHelpMessage = "Usage: boko add|find <args>"
)

func (app *App) addBookmark(args []string) error {

	if len(args) < 2 {
		return errors.New(
			fmt.Sprintf("Not enough arguments.\n%s", addHelpMessage),
		)
	}

	return app.BookmarkRepo.Create(&bookmarks.Bookmark{
		Url:  args[0],
		Name: args[1],
		Tags: args[2:],
	})
}

func (app *App) findBookmarks(args []string) ([]bookmarks.Bookmark, error) {
	if len(args) == 0 {
		return nil, errors.New(fmt.Sprintf("Not enough arguments.\n%s", findHelpMessage))
	}
	if len(args) == 2 {
		switch args[0] {
		case "--name", "-n":
			return app.BookmarkRepo.FindByName(args[1])
		case "--tag", "-t":
			return app.BookmarkRepo.FindByTag(args[1])
		}
	}
	return app.BookmarkRepo.Find(strings.Join(args, " "))
}

func (app *App) Run(args []string) error {
	if len(args) == 0 {
		return errors.New(generalHelpMessage)
	}

	switch args[0] {
	case "add":
		return app.addBookmark(args[1:])
	case "find":
		bookmarks, err := app.findBookmarks(args[1:])
		if err == nil {
			serialized, _ := json.Marshal(bookmarks)
			fmt.Println(string(serialized))
		}
		return err
	default:
		return errors.New(generalHelpMessage)
	}
}
