package handlers

import (
	"github.com/gambitier/goblog/services"
	"github.com/gofiber/fiber/v2"
)

type BlogHandler struct {
	services *services.Service
}

func NewBlogHandler(services *services.Service) *BlogHandler {
	return &BlogHandler{
		services: services,
	}
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
