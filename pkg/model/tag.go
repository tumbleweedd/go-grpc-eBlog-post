package model

type Tag struct {
	Id   int    `json:"-" db:"tag_id"`
	Name string `json:"-" db:"name"`
}
