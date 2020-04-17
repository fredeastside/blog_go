package post

import (
	"blog/pkg/db"
	"fmt"
)

type Posts interface {
	GetBySlug(slug string) (*Post, error)
	GetAll() ([]*Post, error)
}

type PostsRepository struct {
	db *db.DB
	Posts
}

func NewPostsRepository(db *db.DB) Posts {
	return &PostsRepository{
		db: db,
	}
}

func (r *PostsRepository) GetAll() ([]*Post, error) {
	rows, err := r.db.Query("SELECT id, created, name, slug FROM posts ORDER BY created DESC")
	if err != nil {
		return nil, fmt.Errorf("select all error: %s", err.Error())
	}

	var posts []*Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.ID, &post.Created, &post.Name, &post.Slug)
		if err != nil {
			return nil, fmt.Errorf("get all error: %s", err.Error())
		}
		posts = append(posts, &post)
	}

	return posts, nil
}

func (r *PostsRepository) GetBySlug(slug string) (*Post, error) {
	var post Post
	err := r.db.QueryRow("SELECT id, created, name, slug, content FROM posts WHERE slug = ?", slug).
		Scan(&post.ID, &post.Created, &post.Name, &post.Slug, &post.Content)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
