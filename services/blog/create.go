package blog

import (
	"context"

	"github.com/gambitier/goblog/domain/blog"
)

func (s *BlogService) CreateBlog(ctx context.Context, blog *blog.CreateBlogDomainModel) (*blog.Blog, error) {
	params := toCreateDBParams(blog)

	dbBlog, err := s.db.CreateBlogPost(ctx, params)
	if err != nil {
		return nil, err
	}

	return toDomainModel(dbBlog), nil
}
