package repository

import (
	"github.com/jmoiron/sqlx"
	model2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-post/internal/model"
)

type Post interface {
	CreateNewPost(categoryId int, userId int, post model2.PostDTO) error
	GetAllPosts() ([]model2.Post, error)
	GetPostById(id int) (model2.Post, error)
	GetPostsByUserId(userId int) ([]model2.Post, error)
	DeletePostById(postId int) error
}

type Category interface {
	GetCategoryIdByName(categoryName string) (int, error)
	FindCategoryById(categoryId int) (model2.Category, error)
}

type Tag interface {
	GetPostTagsByPostId(postId int) ([]model2.Tag, error)
}

type Repository struct {
	Post
	Category
	Tag
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Post:     NewPostRepository(db),
		Category: NewCategoryRepository(db),
		Tag:      NewTagRepository(db),
	}
}
