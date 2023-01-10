package handler

import (
	"context"

	"github.com/INVITATION-RPC-GATEWAY/api/app/order/pb"
	"github.com/INVITATION-RPC-GATEWAY/api/app/order/services"
	"github.com/gofiber/fiber/v2"
)

var orderservice services.OrderService

type CreateOrderRequestBody struct {
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
	UserId    int64 `json:"userId"`
}

func InitOrder(service services.OrderService) {
	orderservice = service
}

func CreateOrder(ctx *fiber.Ctx) error {
	b := CreateOrderRequestBody{}

	if err := ctx.BodyParser(&b); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": nil})
	}

	res, err := orderservice.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		ProductId: b.ProductId,
		Quantity:  b.Quantity,
		UserId:    b.UserId,
	})

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err, "data": nil})
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "ok", "message": "Success", "data": res})
}
