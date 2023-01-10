package product

import (
	"github.com/INVITATION-RPC-GATEWAY/api/app/product/client"
	"github.com/INVITATION-RPC-GATEWAY/api/app/product/handler"
	"github.com/INVITATION-RPC-GATEWAY/api/app/product/services"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(r fiber.Router) {
	client := client.InitProductClient()

	productService := services.NewProductService(client)
	handler.InitProduct(productService)

	routes := r.Group("/product")
	routes.Post("/", handler.CreateProduct)
	routes.Get("/:id", handler.FindOne)
}