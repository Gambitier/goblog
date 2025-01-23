package middleware

import (
	"errors"

	appErrors "github.com/gambitier/goblog/errors"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Continue stack
		err := c.Next()
		if err == nil {
			return nil
		}

		// Check if it's our custom service error
		var e *appErrors.Error
		if errors.As(err, &e) {
			return c.Status(e.Code).JSON(fiber.Map{
				"error": e.Message,
			})
		}

		// Handle other errors
		return appErrors.NewInternalError(err)
	}
}
