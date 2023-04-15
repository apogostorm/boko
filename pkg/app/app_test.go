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
	testApp.App.AddBookmark([]string{"banana.com", "banana"})
}

func TestAddRequireNameAndUrl(t *testing.T) {
	testApp := getTestApp(t)

	err := testApp.App.AddBookmark([]string{"banana.com"})

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

	testApp.App.AddBookmark([]string{"", "", "fruit", "yellow"})
}

func TestAddReturnsDBError(t *testing.T) {
	testApp := getTestApp(t)
	testApp.RepoMock.
		EXPECT().
		Create(gomock.Any()).
		Return(fmt.Errorf("DB error"))

	err := testApp.App.AddBookmark([]string{"", ""})

	if err == nil {
		t.Errorf("Expected an error when saving to db")
	}
}
