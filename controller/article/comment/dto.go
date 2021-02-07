package comment

import (
	"micro_blog/controller/base"
	"time"
)

type DtoElement struct {
	ID        uint64    `json:"id"`
	Body      string    `json:"body"`
	CommentID *uint64   `json:"comment_id"`
	CreatedAt time.Time `json:"created_at"`
	User      base.User `json:"user"`
}

type Request struct {
	Body      string  `json:"body"`
	CommentID *uint64 `json:"comment_id"`
}
