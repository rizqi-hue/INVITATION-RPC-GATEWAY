package handler

import (
	"context"

	"github.com/INVITATION-RPC-GATEWAY/api/app/auth/pb"
	"github.com/INVITATION-RPC-GATEWAY/api/app/auth/services"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/status"
)

var authservice services.AuthService

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func InitAuth(service services.AuthService) {
	authservice = service
}

func Register(ctx *fiber.Ctx) error {
	b := RegisterRequestBody{}

	if err := ctx.BodyParser(&b); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": nil})
	}

	res, err := authservice.Register(context.Background(), &pb.RegisterRequest{
		Name:     b.Name,
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		errStatus, _ := status.FromError(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": errStatus.Message(), "data": res})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok", "message": "Success", "data": res})
}

func Login(ctx *fiber.Ctx) error {
	b := LoginRequestBody{}

	if err := ctx.BodyParser(&b); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": nil})
	}

	res, err := authservice.Login(context.Background(), &pb.LoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		errStatus, _ := status.FromError(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": errStatus.Message(), "data": res})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok", "message": "Success", "data": res})
}
