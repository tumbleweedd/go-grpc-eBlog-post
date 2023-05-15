package model

type Category struct {
	Id   int    `json:"id" db:"category_id"`
	Name string `json:"name" db:"name"`
}
