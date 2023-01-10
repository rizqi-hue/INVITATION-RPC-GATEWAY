package api

import (
	"fmt"

	apps "github.com/INVITATION-RPC-GATEWAY/api/app"
	"github.com/INVITATION-RPC-GATEWAY/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewFiberServer() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
        AllowOrigins:     "*",
        AllowCredentials: true,
        AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    }))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Fiber")
	})

	apps.SetupRoutes(app)

	if err := app.Listen(config.GetString("port")); err != nil {
		fmt.Println(err)
	}
}
