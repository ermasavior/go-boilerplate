package model

type Post struct {
	ID      int64  `json:"id,omitempty"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
