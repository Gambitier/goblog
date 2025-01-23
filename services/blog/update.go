package blog

import (
	"context"
	"database/sql"
	"errors"

	database "github.com/gambitier/goblog/database/sqlc"
	"github.com/gambitier/goblog/domain/blog"
)

func (s *BlogService) UpdateBlog(ctx context.Context, id int64, updates *blog.Blog) (*blog.Blog, error) {
	// First check if blog exists
	existing, err := s.db.GetBlogPost(ctx, int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrBlogNotFound
		}
		return nil, err
	}

	// Prepare update params, keeping existing values if update field is empty
	params := database.UpdateBlogPostParams{
		ID:          int32(id),
		Title:       updates.Title,
		Description: sql.NullString{String: updates.Description, Valid: updates.Description != ""},
		Body:        updates.Body,
	}

	// If field is empty, keep existing value
	if updates.Title == "" {
		params.Title = existing.Title
	}
	if updates.Description == "" {
		params.Description = existing.Description
	}
	if updates.Body == "" {
		params.Body = existing.Body
	}

	// Perform update
	dbBlog, err := s.db.UpdateBlogPost(ctx, params)
	if err != nil {
		return nil, err
	}

	return toDomainModel(dbBlog), nil
}
