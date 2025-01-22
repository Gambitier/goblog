package httpserver

import (
	"github.com/gambitier/goblog/api/blogs"
	"github.com/gofiber/fiber/v2"
)

func registerRoutes(app *fiber.App) {
	blogs.RegisterRoutes(app)
}
