package handlers

import (
	"github.com/gambitier/goblog/api/blogs/dto"
	"github.com/gambitier/goblog/domain/blog"
	"github.com/gambitier/goblog/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary Create a blog
// @Description Create a new blog
// @Tags blogs
// @Accept json
// @Produce json
// @Param blog body dto.CreateBlogRequest true "Blog details"
// @Success 200 {object} dto.BlogResponse
// @Router /api/blogs [post]
func (h *BlogHandler) CreateBlog(c *fiber.Ctx) error {
	var req dto.CreateBlogRequest
	if err := c.BodyParser(&req); err != nil {
		return errors.NewBadRequestError("invalid request body")
	}

	blog := &blog.CreateBlogDomainModel{
		Title:       req.Title,
		Description: req.Description,
		Body:        req.Body,
	}

	result, err := h.services.BlogService.CreateBlog(c.Context(), blog)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(dto.BlogResponse{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		Body:        result.Body,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	})
}
