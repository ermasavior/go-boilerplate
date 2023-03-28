package post

import (
	"context"
	"database/sql"
	"go-boilerplate/internal/model"
)

type PostRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) GetPosts(ctx context.Context) ([]model.Post, error) {
	rows, err := r.db.Query(getQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []model.Post

	for rows.Next() {
		var p model.Post
		err = rows.Scan(&p.ID, &p.Title, &p.Content)
		if err != nil {
			break
		}

		posts = append(posts, p)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return posts, nil
}

func (r *PostRepository) CreatePost(ctx context.Context, req model.Post) error {
	_, err := r.db.Exec(insertQuery, req.Title, req.Content)
	if err != nil {
		return err
	}

	return nil
}
