package services

import (
	"context"
	"fmt"

	client "github.com/INVITATION-RPC-GATEWAY/api/app/product/client"
	pb "github.com/INVITATION-RPC-GATEWAY/api/app/product/pb"
)

type ProductService interface {
	CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error)
	FindOne(ctx context.Context, in *pb.FindOneRequest) (*pb.FindOneResponse, error)
	DecreaseStock(ctx context.Context, in *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error)
}

type productService struct {
	productClient client.ProductClient
}

func NewProductService(productClient client.ProductClient) ProductService {
	return &productService{productClient: productClient}
}

func (m *productService) FindOne(ctx context.Context, in *pb.FindOneRequest) (out *pb.FindOneResponse, err error) {

	out, err = m.productClient.ProductService.FindOne(ctx, in)
	if err != nil {
		fmt.Printf("err findOne %s", err.Error())
		return out, err
	}

	return out, err
}

func (m *productService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (out *pb.CreateProductResponse, err error) {

	out, err = m.productClient.ProductService.CreateProduct(ctx, in)
	if err != nil {
		fmt.Printf("err findOne %s", err.Error())
		return out, err
	}

	return out, err
}

func (m *productService) DecreaseStock(ctx context.Context, in *pb.DecreaseStockRequest) (out *pb.DecreaseStockResponse, err error) {

	out, err = m.productClient.ProductService.DecreaseStock(ctx, in)
	if err != nil {
		fmt.Printf("err findOne %s", err.Error())
		return out, err
	}

	return out, err
}
