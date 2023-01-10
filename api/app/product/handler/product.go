package handler

import (
	"context"
	"strconv"

	"github.com/INVITATION-RPC-GATEWAY/api/app/product/pb"
	"github.com/INVITATION-RPC-GATEWAY/api/app/product/services"
	"github.com/gofiber/fiber/v2"
)

var productservice services.ProductService

type CreateProductRequestBody struct {
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}

type FindOneProductRequestBody struct {
	Id int64 `json:"id"`
}

func InitProduct(service services.ProductService) {
	productservice = service
}

func CreateProduct(ctx *fiber.Ctx) error {
	b := CreateProductRequestBody{}

	if err := ctx.BodyParser(&b); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": nil})
	}

	res, err := productservice.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:  b.Name,
		Stock: b.Stock,
		Price: b.Price,
	})

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err, "data": nil})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok", "message": "Success", "data": res})
}

func FindOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	i, _ := strconv.Atoi(id)

	res, err := productservice.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(i),
	})

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err, "data": nil})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok", "message": "Success", "data": res})
}
