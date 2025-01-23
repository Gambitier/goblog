package api

import (
	"github.com/gambitier/goblog/api/blogs"
	"github.com/gambitier/goblog/startup/context"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, ctx *context.AppContext) {
	api := app.Group("/api")

	// Register route groups
	blogs.RegisterRoutes(api, ctx)
	// Add other route groups here
}
