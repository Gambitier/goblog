package blogs

import (
	"github.com/gambitier/goblog/api/blogs/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/blogs", handlers.GetBlogs)
}
