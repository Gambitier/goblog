package blog

import "github.com/gambitier/goblog/errors"

var (
	ErrBlogNotFound = errors.NewNotFoundError("blog not found")
)
