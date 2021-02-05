package article

import "time"

type Dto struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user"`
}

type DtoElement struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	ShortBody string    `json:"short_body"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user"`
}

type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type Request struct {
	Title     string `json:"title"`
	ShortBody string `json:"short_body"`
	Body      string `json:"body"`
}

type PatchRequest struct {
	Title     *string `json:"title"`
	ShortBody *string `json:"short_body"`
	Body      *string `json:"body"`
}

type Response struct {
	ID uint64 `json:"id"`
}
