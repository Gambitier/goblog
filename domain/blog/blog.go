package blog

import "time"

// Blog represents the core domain entity
type Blog struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Body        string    `json:"body"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateBlogDomainModel struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

type PaginationParams struct {
	Page     int
	PageSize int
}

type PaginatedResult struct {
	Data       []*Blog
	Total      int64
	Page       int
	PageSize   int
	TotalPages int
}
