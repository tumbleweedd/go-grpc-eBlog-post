package model

import "time"

type Post struct {
	Id           int       `json:"-" db:"post_id"`
	Body         string    `json:"body" db:"body"`
	DateCreation time.Time `json:"date_creation" db:"date_creation"`
	Head         string    `json:"head" db:"head"`
	CategoryId   int       `json:"-" db:"category_id"`
	UserId       int       `json:"-" db:"user_id"`
}

type PotTag struct {
	PostId int `json:"-" db:"post_id"`
	TagId  int `json:"-" db:"tag_id"`
}

type Comments map[string][]string

type PostDTO struct {
	Body string `json:"body"`
	// DateCreation time.Time `json:"date_creation"`
	Head     string   `json:"title"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
	Comments Comments `json:"comments"`
}
