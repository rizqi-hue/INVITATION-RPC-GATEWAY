package order

import (
	"github.com/INVITATION-RPC-GATEWAY/api/app/order/client"
	"github.com/INVITATION-RPC-GATEWAY/api/app/order/handler"
	"github.com/INVITATION-RPC-GATEWAY/api/app/order/services"
	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(r fiber.Router) {
	client := client.InitOrderClient()

	orderService := services.NewOrderService(client)
	handler.InitOrder(orderService)

	routes := r.Group("/order")
	routes.Post("/", handler.CreateOrder)
}
