package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// DeleteBlog godoc
// @Summary Delete a blog
// @Description Delete a blog post by its ID
// @Tags blogs
// @Accept json
// @Produce json
// @Param id path int true "Blog ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/blogs/{id} [delete]
func (h *BlogHandler) DeleteBlog(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid blog ID",
		})
	}

	err = h.services.BlogService.DeleteBlog(c.Context(), id)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
