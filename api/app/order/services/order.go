package services

import (
	"context"
	"fmt"

	"github.com/INVITATION-RPC-GATEWAY/api/app/order/client"
	"github.com/INVITATION-RPC-GATEWAY/api/app/order/pb"
)

type OrderService interface {
	CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error)
}

type orderService struct {
	orderClient client.OrderClient
}

func NewOrderService(orderClient client.OrderClient) OrderService {
	return &orderService{orderClient: orderClient}
}

func (m *orderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {

	out, err := m.orderClient.OrderService.CreateOrder(ctx, in)
	if err != nil {
		fmt.Printf("err register %s", err.Error())
		return out, err
	}

	return out, err
}
