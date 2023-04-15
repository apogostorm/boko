package bookmarks

import (
	"database/sql"
	"strings"
)

const (
	ensureTableQuery = `
	CREATE TABLE IF NOT EXISTS bookmarks (
		id INTEGER PRIMARY KEY,
		name TEXT,
		url TEXT,
		tags TEXT,
		img_path TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`
	insertQuery = `
	INSERT INTO bookmarks 
		(name, url, tags, img_path)
	VALUES 
		($1, $2, $3, $4)
	RETURNING id
	`
	findByNameQuery = `
	SELECT 
		id, name, url, 
		COALESCE(tags, ''),
		COALESCE(img_path, '') 
	FROM bookmarks WHERE name LIKE ?
	`
	findByTagQuery = `
	SELECT 
		id, name, url, 
		COALESCE(tags, ''),
		COALESCE(img_path, '') 
	FROM bookmarks WHERE tags LIKE ?
	`

	tagsDelimiter = ","
)

//go:generate mockgen -destination=../mocks/mock_repo.go -package=mocks -mock_names Repo=BookmarkRepoMock . Repo
type Repo interface {
	FindByName(string) ([]Bookmark, error)
	FindByTag(string) ([]Bookmark, error)
	Create(b *Bookmark) error
}

type repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repo {
	return &repo{db}
}

func (r *repo) FindByName(query string) ([]Bookmark, error) {
	return r.find(findByNameQuery, query)
}

func (r *repo) FindByTag(query string) ([]Bookmark, error) {
	return r.find(findByTagQuery, query)
}

func (r *repo) Create(b *Bookmark) error {
	if err := r.ensureTable(); err != nil {
		return err
	}

	if err := r.db.QueryRow(insertQuery,
		b.Name,
		b.Url,
		strings.Join(b.Tags, tagsDelimiter),
		b.ImagePath,
	).Scan(&b.Id); err != nil {
		return err
	}

	return nil
}

func (r *repo) find(sqlQuery string, query string) ([]Bookmark, error) {
	if err := r.ensureTable(); err != nil {
		return nil, err
	}

	rows, err := r.db.Query(sqlQuery, "%"+query+"%")
	if err != nil {
		return nil, err
	}

	var bookmarks []Bookmark
	for rows.Next() {
		var b Bookmark
		var tags string
		if err := rows.Scan(&b.Id, &b.Name, &b.Url, &tags, &b.ImagePath); err != nil {
			return nil, err
		}
		b.Tags = strings.Split(tags, tagsDelimiter)
		bookmarks = append(bookmarks, b)
	}

	return bookmarks, nil
}

func (r *repo) ensureTable() error {
	_, err := r.db.Exec(ensureTableQuery)
	return err
}
