package router

import (
	"github.com/INVITATION-RPC-GATEWAY/api/app/auth"
	"github.com/INVITATION-RPC-GATEWAY/api/app/order"
	"github.com/INVITATION-RPC-GATEWAY/api/app/product"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	auth.AuthRoutes(api)
	product.ProductRoutes(api)
	order.OrderRoutes(api)
}
