package dto

import "time"

// CreateBlogRequest represents the incoming request for blog creation
type CreateBlogRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Body        string `json:"body" validate:"required"`
}

// UpdateBlogRequest represents the incoming request for blog update
type UpdateBlogRequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Body        string `json:"body,omitempty"`
}

// BlogResponse represents the API response
type BlogResponse struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Body        string    `json:"body"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// PaginationQuery represents the query parameters for pagination
type PaginationQuery struct {
	Page     int `query:"page" default:"1"`
	PageSize int `query:"page_size" default:"10"`
}

// PaginatedBlogResponse represents the response for paginated blog data
type PaginatedBlogResponse struct {
	Data       []BlogResponse `json:"data"`
	Total      int64          `json:"total"`
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
	TotalPages int            `json:"total_pages"`
}
