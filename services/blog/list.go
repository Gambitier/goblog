package blog

import (
	"context"
	"math"

	sqlcDb "github.com/gambitier/goblog/database/sqlc"
	"github.com/gambitier/goblog/domain/blog"
)

func (s *BlogService) GetBlogs(ctx context.Context, params blog.PaginationParams) (*blog.PaginatedResult, error) {
	offset := (params.Page - 1) * params.PageSize

	// Get paginated blogs
	dbBlogs, err := s.db.ListBlogPosts(ctx, sqlcDb.ListBlogPostsParams{
		Limit:  int32(params.PageSize),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}

	// Get total count
	total, err := s.db.CountBlogPosts(ctx)
	if err != nil {
		return nil, err
	}

	// Convert to domain models
	blogs := make([]*blog.Blog, len(dbBlogs))
	for i, dbBlog := range dbBlogs {
		blogs[i] = toDomainModel(dbBlog)
	}

	totalPages := int(math.Ceil(float64(total) / float64(params.PageSize)))

	return &blog.PaginatedResult{
		Data:       blogs,
		Total:      total,
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalPages: totalPages,
	}, nil
}
