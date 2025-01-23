package handlers

import (
	"github.com/gambitier/goblog/api/blogs/handlers"
	"github.com/gambitier/goblog/services"
)

type Handlers struct {
	BlogHandler *handlers.BlogHandler
}

func NewHandlers(services *services.Service) *Handlers {
	return &Handlers{
		BlogHandler: handlers.NewBlogHandler(services),
	}
}
