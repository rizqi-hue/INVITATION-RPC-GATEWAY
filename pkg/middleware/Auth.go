package middleware

import (
	"context"

	"github.com/INVITATION-RPC-GATEWAY/api/app/auth/pb"
	"github.com/INVITATION-RPC-GATEWAY/api/app/auth/services"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/status"
)

var authservice services.AuthService

func InitMiddleWare(service services.AuthService) {
	authservice = service
}

func Auth(ctx *fiber.Ctx) error {

	h := ctx.Cookies("token")

	if h == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User has been singout", "data": nil})
	}

	res, err := authservice.RefreshToken(context.Background(), &pb.RefreshTokenRequest{
		Token: h,
	})

	if err != nil {
		errStatus, _ := status.FromError(err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": errStatus.Message(), "data": res})
	}

	ctx.Next()
	return nil
}
