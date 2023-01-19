package auth

import (
	"github.com/INVITATION-RPC-GATEWAY/api/app/auth/client"
	"github.com/INVITATION-RPC-GATEWAY/api/app/auth/handler"
	"github.com/INVITATION-RPC-GATEWAY/api/app/auth/services"
	"github.com/INVITATION-RPC-GATEWAY/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(r fiber.Router) {
	client := client.InitAuthClient()

	authService := services.NewAuthService(client)
	handler.InitAuth(authService)
	middleware.InitMiddleWare(authService)

	routes := r.Group("/auth")
	routes.Post("/sign-in", handler.Login)
	routes.Post("/sign-up", handler.Register)
	routes.Post("/sign-out", middleware.Auth, handler.Logout)
	routes.Get("/profile", middleware.Auth, handler.Profile)

	// routes.Post("/refresh-token", middleware.Auth, handler.RefreshToken)
}
