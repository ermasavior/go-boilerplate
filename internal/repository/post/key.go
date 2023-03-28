package post

const (
	getQuery    = "SELECT id, title, content FROM posts"
	insertQuery = "INSERT INTO posts (title, content) VALUES ($1, $2)"
)
