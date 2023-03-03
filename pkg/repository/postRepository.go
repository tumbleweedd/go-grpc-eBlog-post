package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-post/pkg/model"
)

type PostRepository struct {
	db *sqlx.DB
}

func (postRepo *PostRepository) CreateNewPost(categoryId int, userId int, post model.PostDTO) error {
	tx, err := postRepo.db.Begin()
	if err != nil {
		return err
	}

	var postId int
	createPostQuery := fmt.Sprintf(`insert into %s (body, date_creation, head, category_id, user_id)
											values($1, current_timestamp, $2, $3, $4) returning post_id`, postTable)
	row := tx.QueryRow(createPostQuery, post.Body, post.Head, categoryId, userId)
	if err := row.Scan(&postId); err != nil {
		tx.Rollback()
		return err
	}

	for _, tag := range post.Tags {
		// Для тегов нужно сделать проверку: если данный тег уже существует в таблице tag, то мы не добавляем
		// его в таблицу заново, а просто используем. Таким образом, получается избежать холостой вставки
		// в таблицу tag
		var tagId int
		createTagQuery := fmt.Sprintf(`insert into %s (name) values($1) returning tag_id`, tagTable)
		row := tx.QueryRow(createTagQuery, tag)
		if err := row.Scan(&tagId); err != nil {
			tx.Rollback()
			return err
		}

		createPostTagQuery := fmt.Sprintf(`insert into %s (post_id, tag_id) values($1, $2)`, postTagTable)
		_, err := tx.Exec(createPostTagQuery, postId, tagId)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (postRepo *PostRepository) GetAllPosts() ([]model.Post, error) {
	var posts []model.Post

	getAllPostsQuery := fmt.Sprintf(`select * from %s p order by date_creation desc`, postTable)

	err := postRepo.db.Select(&posts, getAllPostsQuery)

	return posts, err
}

func (postRepo *PostRepository) GetPostsByUserId(userId int) ([]model.Post, error) {
	var userPosts []model.Post

	getUserPostsQuery := fmt.Sprintf(`select * from %s p where p.user_id=$1 order by date_creation desc`, postTable)

	err := postRepo.db.Select(&userPosts, getUserPostsQuery, userId)

	return userPosts, err
}

func (postRepo *PostRepository) GetPostById(postId int) (model.Post, error) {
	var post model.Post

	getPostQuery := fmt.Sprintf(`select * from %s p where p.post_id = $1`, postTable)

	err := postRepo.db.Get(&post, getPostQuery, postId)

	return post, err
}

func (postRepo *PostRepository) GetAllPostsByUserId(userId int) ([]model.Post, error) {
	var posts []model.Post

	getAllPostsQuery := fmt.Sprintf(`select * from %s p where p.user_id = $1 order by date_creation desc`, postTable)

	err := postRepo.db.Select(&posts, getAllPostsQuery, userId)

	return posts, err
}

func NewPostRepository(db *sqlx.DB) *PostRepository {
	return &PostRepository{db: db}
}
