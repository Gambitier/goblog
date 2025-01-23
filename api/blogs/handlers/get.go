package handlers

import (
	"strconv"

	"github.com/gambitier/goblog/api/blogs/dto"
	"github.com/gofiber/fiber/v2"
)

// GetBlog godoc
// @Summary Get a single blog
// @Description Get a blog post by its ID
// @Tags blogs
// @Accept json
// @Produce json
// @Param id path int true "Blog ID"
// @Success 200 {object} dto.BlogResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/blogs/{id} [get]
func (h *BlogHandler) GetBlog(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid blog ID",
		})
	}

	blog, err := h.services.BlogService.GetBlog(c.Context(), int32(id))
	if err != nil {
		return err
	}

	return c.JSON(dto.BlogResponse{
		ID:          blog.ID,
		Title:       blog.Title,
		Description: blog.Description,
		Body:        blog.Body,
		CreatedAt:   blog.CreatedAt,
		UpdatedAt:   blog.UpdatedAt,
	})
}
