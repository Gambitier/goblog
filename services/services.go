package services

import (
	blogservice "github.com/gambitier/goblog/services/blog"
)

type Service struct {
	BlogService *blogservice.BlogService
}

func NewService() (*Service, error) {
	blogService := blogservice.NewBlogService()
	return &Service{
		BlogService: blogService,
	}, nil
}
