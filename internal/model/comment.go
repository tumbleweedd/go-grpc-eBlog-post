package model

import "time"

type Comment struct {
	Id           int       `json:"id" db:"comment_id"`
	Body         string    `json:"body" db:"body"`
	DateCreation time.Time `json:"dateCreation" db:"date_creation"`
	PostId       int       `json:"-" db:"post_id"`
	UserId       int       `json:"-" db:"user_id"`
}
