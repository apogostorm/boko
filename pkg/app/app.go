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
	helpMessage = "Usage: boko add <url> <name> [tags...]"
)

func (app *App) AddBookmark(args []string) error {

	if len(args) < 2 {
		return errors.New(
			fmt.Sprintf("Not enough arguments.\n%s", helpMessage),
		)
	}

	return app.BookmarkRepo.Create(&bookmarks.Bookmark{
		Url:  args[0],
		Name: args[1],
		Tags: args[2:],
	})
}
