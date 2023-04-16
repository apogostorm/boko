package app

import (
	"fmt"
	"testing"

	"github.com/apogostorm/boko/pkg/bookmarks"
	"github.com/apogostorm/boko/pkg/mocks"
	"github.com/golang/mock/gomock"
)

type testApp struct {
	RepoMock *mocks.BookmarkRepoMock
	App      *App
}

func getTestApp(t *testing.T) *testApp {
	ctrl := gomock.NewController(t)
	repo := mocks.NewBookmarkRepoMock(ctrl)

	return &testApp{
		RepoMock: repo,
		App: &App{
			BookmarkRepo: repo,
		},
	}
}

func TestAddSavesToDb(t *testing.T) {
	testApp := getTestApp(t)
	testApp.RepoMock.
		EXPECT().
		Create(gomock.Eq(&bookmarks.Bookmark{
			Url:  "banana.com",
			Name: "banana",
			Tags: []string{},
		}))
	testApp.App.Run([]string{"add", "banana.com", "banana"})
}

func TestAddRequireNameAndUrl(t *testing.T) {
	testApp := getTestApp(t)

	err := testApp.App.Run([]string{"add", "banana.com"})

	if err == nil {
		t.Errorf("Expected an error when giving only a url")
	}
	fmt.Println(err)
}

func TestAddSavesTags(t *testing.T) {
	testApp := getTestApp(t)
	testApp.RepoMock.
		EXPECT().
		Create(gomock.Eq(&bookmarks.Bookmark{
			Tags: []string{"fruit", "yellow"},
		}))

	testApp.App.Run([]string{"add", "", "", "fruit", "yellow"})
}

func TestAddReturnsDBError(t *testing.T) {
	testApp := getTestApp(t)
	testApp.RepoMock.
		EXPECT().
		Create(gomock.Any()).
		Return(fmt.Errorf("DB error"))

	err := testApp.App.Run([]string{"add", "", ""})

	if err == nil {
		t.Errorf("Expected an error when saving to db")
	}
}

func TestRundFindCallsFindBookmark(t *testing.T) {
	testApp := getTestApp(t)
	testApp.RepoMock.
		EXPECT().
		FindByName("banana").
		Return([]bookmarks.Bookmark{{Id: 123}}, nil)

	testApp.App.Run([]string{"find", "-n", "banana"})
}

func findByName(t *testing.T, option string) {
	testApp := getTestApp(t)
	testApp.RepoMock.
		EXPECT().
		FindByName("banana").
		Return([]bookmarks.Bookmark{{Id: 123}}, nil)

	items, err := testApp.App.findBookmarks([]string{option, "banana"})

	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if len(items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(items))
	}
	if items[0].Id != 123 {
		t.Errorf("Expected item with id 123, got %d", items[0].Id)
	}
}

func TestFindByNameLongForm(t *testing.T) {
	findByName(t, "--name")
}

func TestFindByNameShortForm(t *testing.T) {
	findByName(t, "-n")
}

func TestFindByNameErrorWhenNotEnoughArgs(t *testing.T) {
	testApp := getTestApp(t)

	_, err := testApp.App.findBookmarks([]string{"-n"})

	if err == nil {
		t.Errorf("Expected an error when not giving enough arguments")
	}
}

func TestFindBookmarksErrorWhenNotEnoughArgs(t *testing.T) {
	testApp := getTestApp(t)

	_, err := testApp.App.findBookmarks([]string{})

	if err == nil {
		t.Errorf("Expected an error when not giving enough arguments")
	}
}

func testFindByTag(t *testing.T, option string) {
	testApp := getTestApp(t)
	testApp.RepoMock.
		EXPECT().
		FindByTag("banana").
		Return([]bookmarks.Bookmark{{Id: 123}}, nil)

	items, err := testApp.App.findBookmarks([]string{option, "banana"})

	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if len(items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(items))
	}
	if items[0].Id != 123 {
		t.Errorf("Expected item with id 123, got %d", items[0].Id)
	}
}

func TestFindByTagLongForm(t *testing.T) {
	testFindByTag(t, "--tag")
}

func TestFindByTagShortForm(t *testing.T) {
	testFindByTag(t, "-t")
}

func TestFindByTagErrorWhenNotEnoughArgs(t *testing.T) {
	testApp := getTestApp(t)

	_, err := testApp.App.findBookmarks([]string{"-t"})

	if err == nil {
		t.Errorf("Expected an error when not giving enough arguments")
	}
}
