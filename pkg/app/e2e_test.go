package app

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/apogostorm/boko/pkg/bookmarks"
	_ "github.com/mattn/go-sqlite3"
)

var data []bookmarks.Bookmark = []bookmarks.Bookmark{{
	Name: "banana",
	Url:  "banana.com",
	Tags: []string{"fruit", "yellow"},
}, {
	Name: "apple",
	Url:  "apple.com",
	Tags: []string{"fruit", "red"},
}, {
	Name: "carrot",
	Url:  "carrot.com",
	Tags: []string{"vegetable", "orange"},
}}

func setup(t *testing.T) (App, func(t *testing.T)) {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("sqlite3", dir+"/test.db")
	if err != nil {
		log.Fatal(err)
	}
	repo := bookmarks.NewRepo(db)

	return App{BookmarkRepo: repo}, func(t *testing.T) {
		os.RemoveAll(dir)
		db.Close()
	}
}

func TestAppE2e(t *testing.T) {
	app, teardown := setup(t)
	defer teardown(t)

	for _, d := range data {
		app.Run([]string{"add", d.Url, d.Name, d.Tags[0], d.Tags[1]})
	}
	// Find by name
	e2eTestFindByName(t, app)
	e2eTestFindByTag(t, app)
}

func e2eTestFindByName(t *testing.T, app App) {
	bookmarks, err := app.findBookmarks([]string{"-n", "banana"})
	if err != nil {
		t.Errorf("Error while finding bookmark: %s", err)
	}

	if len(bookmarks) != 1 {
		t.Errorf("Expected 1 bookmark, got %d", len(bookmarks))
	}

	bookmark := bookmarks[0]
	d := data[0]

	if bookmark.Name != d.Name {
		t.Errorf("Expected name %s, got %s", d.Name, bookmark.Name)
	}
	if bookmark.Url != d.Url {
		t.Errorf("Expected url %s, got %s", d.Url, bookmark.Url)
	}
	if bookmark.Tags[0] != d.Tags[0] {
		t.Errorf("Expected tag %s, got %s", d.Tags[0], bookmark.Tags[0])
	}
	if bookmark.Tags[1] != d.Tags[1] {
		t.Errorf("Expected tag %s, got %s", d.Tags[1], bookmark.Tags[1])
	}
}

func e2eTestFindByTag(t *testing.T, app App) {
	bookmarks, err := app.findBookmarks([]string{"-t", "fruit"})
	if err != nil {
		t.Errorf("Error while finding bookmark: %s", err)
	}

	if len(bookmarks) != 2 {
		t.Errorf("Expected 2 bookmark, got %d", len(bookmarks))
	}

	for i, bookmark := range bookmarks {

		if bookmark.Name != data[i].Name {
			t.Errorf("Expected name %s, got %s", data[i].Name, bookmark.Name)
		}
		if bookmark.Url != data[i].Url {
			t.Errorf("Expected url %s, got %s", data[i].Url, bookmark.Url)
		}
		if bookmark.Tags[0] != data[i].Tags[0] {
			t.Errorf("Expected tag %s, got %s", data[i].Tags[0], bookmark.Tags[0])
		}
		if bookmark.Tags[1] != data[i].Tags[1] {
			t.Errorf("Expected tag %s, got %s", data[i].Tags[1], bookmark.Tags[1])
		}

	}
}
