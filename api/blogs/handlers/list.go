package handlers

import (
	"github.com/gambitier/goblog/api/blogs/dto"
	"github.com/gambitier/goblog/domain/blog"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get blogs
// @Description Get all blogs with pagination
// @Tags blogs
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)" minimum(1)
// @Param page_size query int false "Items per page (default: 10)" minimum(1)
// @Success 200 {object} dto.PaginatedBlogResponse
// @Router /api/blogs [get]
func (h *BlogHandler) GetBlogs(c *fiber.Ctx) error {
	var query dto.PaginationQuery
	if err := c.QueryParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid query parameters",
		})
	}

	// Set defaults if not provided
	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 {
		query.PageSize = 10
	}

	result, err := h.services.BlogService.GetBlogs(c.Context(), blog.PaginationParams{
		Page:     query.Page,
		PageSize: query.PageSize,
	})
	if err != nil {
		return err
	}

	// Convert domain models to response DTOs
	blogs := make([]dto.BlogResponse, len(result.Data))
	for i, blog := range result.Data {
		blogs[i] = dto.BlogResponse{
			ID:          blog.ID,
			Title:       blog.Title,
			Description: blog.Description,
			Body:        blog.Body,
			CreatedAt:   blog.CreatedAt,
			UpdatedAt:   blog.UpdatedAt,
		}
	}

	return c.JSON(dto.PaginatedBlogResponse{
		Data:       blogs,
		Total:      result.Total,
		Page:       result.Page,
		PageSize:   result.PageSize,
		TotalPages: result.TotalPages,
	})
}
