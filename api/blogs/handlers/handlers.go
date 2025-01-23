package handlers

import (
	"github.com/gambitier/goblog/services"
)

type BlogHandler struct {
	services *services.Service
}

func NewBlogHandler(services *services.Service) *BlogHandler {
	return &BlogHandler{
		services: services,
	}
}
