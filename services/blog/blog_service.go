package blog

import (
	"github.com/gambitier/goblog/database"
)

type BlogService struct {
	db *database.Database
}

func NewBlogService(db *database.Database) *BlogService {
	return &BlogService{db: db}
}
