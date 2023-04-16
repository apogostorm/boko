package app

import (
	"errors"
	"fmt"

	"github.com/apogostorm/boko/pkg/bookmarks"
)

type App struct {
	BookmarkRepo bookmarks.Repo
}

const (
	addHelpMessage     = "Usage: boko add <url> <name> [tags...]"
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

func (app *App) Run(args []string) error {
	if len(args) == 0 {
		return errors.New(generalHelpMessage)
	}

	switch args[0] {
	case "add":
		return app.addBookmark(args[1:])
	default:
		return errors.New(generalHelpMessage)
	}
}
