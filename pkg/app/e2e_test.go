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
	e2eTestFind(t, app)
}

func e2eTestFindByName(t *testing.T, app App) {
	results, err := app.findBookmarks([]string{"-n", "banana"})
	if err != nil {
		t.Errorf("Error while finding bookmark: %s", err)
	}
	compareBookmarks(t, data[0:1], results)
}

func e2eTestFindByTag(t *testing.T, app App) {
	results, err := app.findBookmarks([]string{"-t", "fruit"})
	if err != nil {
		t.Errorf("Error while finding bookmark: %s", err)
	}
	compareBookmarks(t, data[0:2], results)
}

func e2eTestFind(t *testing.T, app App) {
	results, err := app.findBookmarks([]string{"an"})
	if err != nil {
		t.Errorf("Error while finding bookmark: %s", err)
	}
	compareBookmarks(t, []bookmarks.Bookmark{data[0], data[2]}, results)
}

func compareBookmarks(t *testing.T, expected []bookmarks.Bookmark, results []bookmarks.Bookmark) {
	if len(results) != len(expected) {
		t.Errorf("Expected %d bookmark(s), got %d", len(expected), len(results))
	}
	for i, result := range results {
		if result.Name != expected[i].Name {
			t.Errorf("Expected name %s, got %s", expected[i].Name, result.Name)
		}
		if result.Url != expected[i].Url {
			t.Errorf("Expected url %s, got %s", expected[i].Url, result.Url)
		}
		if result.Tags[0] != expected[i].Tags[0] {
			t.Errorf("Expected tag %s, got %s", expected[i].Tags[0], result.Tags[0])
		}
		if result.Tags[1] != expected[i].Tags[1] {
			t.Errorf("Expected tag %s, got %s", expected[i].Tags[1], result.Tags[1])
		}
	}
}
