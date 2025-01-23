package blogs

import (
	"github.com/gambitier/goblog/startup/context"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, ctx *context.AppContext) {
	blogs := router.Group("/blogs")

	blogs.Get("/", ctx.Handlers.BlogHandler.GetBlogs)
	blogs.Post("/", ctx.Handlers.BlogHandler.CreateBlog)
	blogs.Get("/:id", ctx.Handlers.BlogHandler.GetBlog)
	blogs.Patch("/:id", ctx.Handlers.BlogHandler.UpdateBlog)
	blogs.Delete("/:id", ctx.Handlers.BlogHandler.DeleteBlog)
}
