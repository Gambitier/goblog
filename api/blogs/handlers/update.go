package handlers

import (
	"strconv"

	"github.com/gambitier/goblog/api/blogs/dto"
	"github.com/gambitier/goblog/domain/blog"
	"github.com/gambitier/goblog/errors"
	"github.com/gofiber/fiber/v2"
)

// UpdateBlog godoc
// @Summary Update a blog
// @Description Update a blog post by its ID
// @Tags blogs
// @Accept json
// @Produce json
// @Param id path int true "Blog ID"
// @Param blog body dto.UpdateBlogRequest true "Blog update object"
// @Success 200 {object} dto.BlogResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/blogs/{id} [put]
func (h *BlogHandler) UpdateBlog(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return errors.NewBadRequestError("invalid blog ID")
	}

	var req dto.UpdateBlogRequest
	if err := c.BodyParser(&req); err != nil {
		return errors.NewBadRequestError("invalid request body")
	}

	blog, err := h.services.BlogService.UpdateBlog(c.Context(), id, &blog.Blog{
		Title:       req.Title,
		Description: req.Description,
		Body:        req.Body,
	})
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
