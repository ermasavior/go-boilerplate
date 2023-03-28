package post

import (
	"context"
	"database/sql"
	"errors"
	"go-boilerplate/internal/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func initMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestRepository_GetPosts(t *testing.T) {
	db, mock := initMock(t)
	defer db.Close()

	repo := New(db)

	t.Run("success - return posts", func(t *testing.T) {
		expected := []model.Post{
			{
				ID:      1,
				Title:   "hello",
				Content: "hello world",
			},
		}

		rows := sqlmock.NewRows([]string{"id", "title", "content"})
		rows.AddRow(1, "hello", "hello world")
		mock.ExpectQuery("SELECT id, title, content FROM posts").WillReturnRows(rows)

		got, err := repo.GetPosts(context.TODO())

		assert.Equal(t, expected, got)
		assert.Equal(t, true, err == nil)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %s", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, title, content FROM posts").WillReturnError(errors.New("error"))

		_, err := repo.GetPosts(context.TODO())

		assert.Equal(t, false, err == nil)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %s", err)
		}
	})
}

func TestRepository_CreatePost(t *testing.T) {
	db, mock := initMock(t)
	defer db.Close()

	repo := New(db)

	req := model.Post{
		Title:   "hello",
		Content: "hello world",
	}

	t.Run("success - return posts", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO posts (.+) VALUES (.+)").
			WithArgs("hello", "hello world").
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.CreatePost(context.TODO(), req)

		assert.Equal(t, true, err == nil)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %s", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO posts (.+) VALUES (.+)").
			WithArgs("hello", "hello world").
			WillReturnError(errors.New("error"))

		err := repo.CreatePost(context.TODO(), req)

		assert.Equal(t, false, err == nil)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %s", err)
		}
	})
}
