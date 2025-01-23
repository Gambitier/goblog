package blog

import (
	"database/sql"

	sqlcDb "github.com/gambitier/goblog/database/sqlc"
	"github.com/gambitier/goblog/domain/blog"
)

func toCreateDBParams(blog *blog.CreateBlogDomainModel) sqlcDb.CreateBlogPostParams {
	return sqlcDb.CreateBlogPostParams{
		Title:       blog.Title,
		Description: sql.NullString{String: blog.Description, Valid: blog.Description != ""},
		Body:        blog.Body,
	}
}

func toDomainModel(dbBlog sqlcDb.BlogPost) *blog.Blog {
	return &blog.Blog{
		ID:          int64(dbBlog.ID),
		Title:       dbBlog.Title,
		Description: dbBlog.Description.String,
		Body:        dbBlog.Body,
		CreatedAt:   dbBlog.CreatedAt.Time,
		UpdatedAt:   dbBlog.UpdatedAt.Time,
	}
}
