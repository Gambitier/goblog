package blog

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gambitier/goblog/domain/blog"
)

func (s *BlogService) GetBlog(ctx context.Context, id int32) (*blog.Blog, error) {
	dbBlog, err := s.db.GetBlogPost(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrBlogNotFound
		}
		return nil, err
	}

	return toDomainModel(dbBlog), nil
}
