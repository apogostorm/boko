// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/apogostorm/boko/pkg/bookmarks (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	bookmarks "github.com/apogostorm/boko/pkg/bookmarks"
	gomock "github.com/golang/mock/gomock"
)

// BookmarkRepoMock is a mock of Repo interface.
type BookmarkRepoMock struct {
	ctrl     *gomock.Controller
	recorder *BookmarkRepoMockMockRecorder
}

// BookmarkRepoMockMockRecorder is the mock recorder for BookmarkRepoMock.
type BookmarkRepoMockMockRecorder struct {
	mock *BookmarkRepoMock
}

// NewBookmarkRepoMock creates a new mock instance.
func NewBookmarkRepoMock(ctrl *gomock.Controller) *BookmarkRepoMock {
	mock := &BookmarkRepoMock{ctrl: ctrl}
	mock.recorder = &BookmarkRepoMockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *BookmarkRepoMock) EXPECT() *BookmarkRepoMockMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *BookmarkRepoMock) Create(arg0 *bookmarks.Bookmark) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *BookmarkRepoMockMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*BookmarkRepoMock)(nil).Create), arg0)
}

// Find mocks base method.
func (m *BookmarkRepoMock) Find(arg0 string) ([]bookmarks.Bookmark, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].([]bookmarks.Bookmark)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *BookmarkRepoMockMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*BookmarkRepoMock)(nil).Find), arg0)
}

// FindByName mocks base method.
func (m *BookmarkRepoMock) FindByName(arg0 string) ([]bookmarks.Bookmark, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", arg0)
	ret0, _ := ret[0].([]bookmarks.Bookmark)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *BookmarkRepoMockMockRecorder) FindByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*BookmarkRepoMock)(nil).FindByName), arg0)
}

// FindByTag mocks base method.
func (m *BookmarkRepoMock) FindByTag(arg0 string) ([]bookmarks.Bookmark, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByTag", arg0)
	ret0, _ := ret[0].([]bookmarks.Bookmark)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTag indicates an expected call of FindByTag.
func (mr *BookmarkRepoMockMockRecorder) FindByTag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTag", reflect.TypeOf((*BookmarkRepoMock)(nil).FindByTag), arg0)
}
