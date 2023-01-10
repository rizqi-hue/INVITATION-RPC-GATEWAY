package auth

import (
	"github.com/INVITATION-RPC-GATEWAY/api/app/auth/client"
	"github.com/INVITATION-RPC-GATEWAY/api/app/auth/handler"
	"github.com/INVITATION-RPC-GATEWAY/api/app/auth/services"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(r fiber.Router) {
	client := client.InitAuthClient()

	authService := services.NewAuthService(client)
	handler.InitAuth(authService)

	routes := r.Group("/auth")
	routes.Post("/sign-in", handler.Login)
	routes.Post("/sign-up", handler.Register)
}
