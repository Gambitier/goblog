package services

import (
	"github.com/gambitier/goblog/database"
	blogservice "github.com/gambitier/goblog/services/blog"
)

type Service struct {
	BlogService *blogservice.BlogService
}

func NewService(db *database.Database) (*Service, error) {
	blogService := blogservice.NewBlogService(db)
	return &Service{
		BlogService: blogService,
	}, nil
}
