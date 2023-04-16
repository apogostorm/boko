package bookmarks

type Bookmark struct {
	Id        int64    `json:"id"`
	Name      string   `json:"name"`
	Url       string   `json:"url"`
	Tags      []string `json:"tags"`
	ImagePath string   `json:"imagePath"`
}
