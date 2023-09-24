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

var bookmarksResults = []bookmarks.Bookmark{{Id: 123}}

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
			Url:       "banana.com",
			Name:      "banana",
			Tags:      []string{},
			ImagePath: "~/.boko/icons/default.png",
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
			ImagePath: "~/.boko/icons/default.png",
			Tags:      []string{"fruit", "yellow"},
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

func testFindByName(t *testing.T, option string) {
	testApp := getTestApp(t)
	testApp.RepoMock.
		EXPECT().
		FindByName("banana").
		Return(bookmarksResults, nil)

	items, err := testApp.App.findBookmarks([]string{option, "banana"})

	nilError(t, err)
	assertExpectedResults(t, items)
}

func TestFindByNameLongForm(t *testing.T) {
	testFindByName(t, "--name")
}

func TestFindByNameShortForm(t *testing.T) {
	testFindByName(t, "-n")
}

func testFindByTag(t *testing.T, option string) {
	testApp := getTestApp(t)
	testApp.RepoMock.
		EXPECT().
		FindByTag("banana").
		Return(bookmarksResults, nil)

	items, err := testApp.App.findBookmarks([]string{option, "banana"})

	nilError(t, err)
	assertExpectedResults(t, items)
}

func TestFindByTagLongForm(t *testing.T) {
	testFindByTag(t, "--tag")
}

func TestFindByTagShortForm(t *testing.T) {
	testFindByTag(t, "-t")
}

func TestFindOneArgument(t *testing.T) {
	testApp := getTestApp(t)
	testApp.RepoMock.
		EXPECT().
		Find("query").
		Return(bookmarksResults, nil)
	items, err := testApp.App.findBookmarks([]string{"query"})

	nilError(t, err)
	assertExpectedResults(t, items)
}

func TestFindManyArguments(t *testing.T) {
	testApp := getTestApp(t)
	testApp.RepoMock.
		EXPECT().
		Find("query that is long").
		Return(bookmarksResults, nil)
	_, err := testApp.App.findBookmarks([]string{"query", "that", "is", "long"})

	nilError(t, err)
}

func nilError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}

func assertExpectedResults(t *testing.T, results []bookmarks.Bookmark) {
	if len(results) != 1 {
		t.Errorf("Expected 1 item, got %d", len(results))
	}
	if results[0].Id != 123 {
		t.Errorf("Expected item with id 123, got %d", results[0].Id)
	}
}
