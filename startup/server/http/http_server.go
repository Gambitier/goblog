package httpserver

import (
	"fmt"
	"log"

	"github.com/gambitier/goblog/api"
	"github.com/gambitier/goblog/middleware"
	"github.com/gambitier/goblog/startup/context"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type HttpServer struct {
	app        *fiber.App
	appContext *context.AppContext
}

func NewHttpServer(appContext *context.AppContext) *HttpServer {
	return &HttpServer{
		app:        fiber.New(),
		appContext: appContext,
	}
}

func (server *HttpServer) Configure() {
	server.app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: false,
	}))

	server.app.Use(etag.New())
	server.app.Use(cache.New())
	server.app.Use(compress.New())
	server.app.Use(recover.New())
	server.app.Use(idempotency.New())
	server.app.Use(func(c *fiber.Ctx) error {
		if c.Path() == "/swagger" {
			return c.Next()
		}
		return helmet.New()(c)
	})

	server.app.Use(logger.New())

	server.app.Use(swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: "./_apidocs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	}))

	server.app.Use(middleware.ErrorHandler())
}

func (server *HttpServer) RunServer() error {
	address := fmt.Sprintf(":%v", server.appContext.Config.Port)
	log.Printf(
		"Starting server on PORT: %v & Host: %s\n",
		server.appContext.Config.Port,
		server.appContext.Config.APIHost,
	)

	if err := server.app.Listen(address); err != nil {
		return err
	}

	return nil
}

func (server *HttpServer) Shutdown() error {
	return server.app.Shutdown()
}

func (server *HttpServer) RegisterRoutes() {
	api.SetupRoutes(server.app, server.appContext)
}
