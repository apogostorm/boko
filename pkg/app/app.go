package app

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/apogostorm/boko/pkg/bookmarks"
)

type App struct {
	BookmarkRepo bookmarks.Repo
}

const (
	addHelpMessage     = "Usage: boko add <url> <name> [tags...]"
	findHelpMessage    = "Usage: boko find --name|-n <name>"
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
	switch args[0] {
	case "--name", "-n":
		if len(args) < 2 {
			return nil, errors.New(fmt.Sprintf("Not enough arguments.\n%s", findHelpMessage))
		}
		return app.BookmarkRepo.FindByName(args[1])
	case "--tag", "-t":
		if len(args) < 2 {
			return nil, errors.New(fmt.Sprintf("Not enough arguments.\n%s", findHelpMessage))
		}
		return app.BookmarkRepo.FindByTag(args[1])
	default:
		return nil, errors.New("not implemented")
	}
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
