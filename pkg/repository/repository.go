package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-post/pkg/model"
)

type Post interface {
	CreateNewPost(categoryId int, userId int, post model.PostDTO) error
	GetAllPosts() ([]model.Post, error)
	GetPostById(id int) (model.Post, error)
	GetPostsByUserId(userId int) ([]model.Post, error)
	DeletePostById(postId int) error
}

type Category interface {
	GetCategoryIdByName(categoryName string) (int, error)
	FindCategoryById(categoryId int) (model.Category, error)
}

type Tag interface {
	GetPostTagsByPostId(postId int) ([]model.Tag, error)
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
