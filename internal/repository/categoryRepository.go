package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-post/internal/model"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (categoryRepo *CategoryRepository) GetCategoryIdByName(categoryName string) (int, error) {
	var category model.Category

	getCategoryIdQuery := fmt.Sprintf(`select * from %s c where c.name = $1`, categoryTable)
	err := categoryRepo.db.Get(&category, getCategoryIdQuery, categoryName)

	return category.Id, err
}

func (categoryRepo *CategoryRepository) FindCategoryById(categoryId int) (model.Category, error) {
	var category model.Category

	getCategoryIdQuery := fmt.Sprintf(`select * from %s c where c.category_id = $1`, categoryTable)
	err := categoryRepo.db.Get(&category, getCategoryIdQuery, categoryId)

	return category, err
}
