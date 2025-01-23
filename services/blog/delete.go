package blog

import (
	"context"
	"database/sql"
	"errors"
)

func (s *BlogService) DeleteBlog(ctx context.Context, id int64) error {
	// First check if blog exists
	_, err := s.db.GetBlogPost(ctx, int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrBlogNotFound
		}
		return err
	}

	// Delete the blog
	err = s.db.DeleteBlogPost(ctx, int32(id))
	if err != nil {
		return err
	}

	return nil
}
