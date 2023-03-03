package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-post/pkg/model"
)

type TagRepository struct {
	db *sqlx.DB
}

func NewTagRepository(db *sqlx.DB) *TagRepository {
	return &TagRepository{db: db}
}

func (tagRepo *TagRepository) GetPostTagsByPostId(postId int) ([]model.Tag, error) {
	var tags []model.Tag

	getPostTagsQuery := fmt.Sprintf(`select * from %s t where t.tag_id in 
									(select pt.tag_id from %s pt where pt.post_id = $1)`,
		tagTable, postTagTable)
	err := tagRepo.db.Select(&tags, getPostTagsQuery, postId)

	return tags, err

}
