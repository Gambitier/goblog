package handlers

import (
	"github.com/gambitier/goblog/services"
	"github.com/gofiber/fiber/v2"
)

type Blog struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type BlogHandler struct {
	services *services.Service
}

func NewBlogHandler(services *services.Service) *BlogHandler {
	return &BlogHandler{
		services: services,
	}
}

// @Summary Get blogs
// @Description Get all blogs
// @Tags blogs
// @Accept json
// @Produce json
// @Success 200 {array} Blog
// @Router /blogs [get]
func (h *BlogHandler) GetBlogs(c *fiber.Ctx) error {
	// Use h.services.BlogService to interact with the database
	return c.JSON(nil)
}

func (h *BlogHandler) CreateBlog(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (h *BlogHandler) GetBlog(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (h *BlogHandler) UpdateBlog(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (h *BlogHandler) DeleteBlog(c *fiber.Ctx) error {
	return c.JSON(nil)
}
