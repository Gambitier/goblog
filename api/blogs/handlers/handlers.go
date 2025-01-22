package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type Blog struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// @Summary Get blogs
// @Description Get all blogs
// @Tags blogs
// @Accept json
// @Produce json
// @Success 200 {array} Blog
// @Router /blogs [get]
func GetBlogs(c *fiber.Ctx) error {
	blogs := []Blog{
		{ID: "1", Title: "Blog 1", Content: "Content 1", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
		{ID: "2", Title: "Blog 2", Content: "Content 2", CreatedAt: "2021-01-02", UpdatedAt: "2021-01-02"},
	}
	return c.JSON(blogs)
}
