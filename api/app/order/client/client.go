package client

import (
	"github.com/INVITATION-RPC-GATEWAY/api/app/order/pb"
	"github.com/INVITATION-RPC-GATEWAY/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderClient struct {
	OrderService pb.OrderServiceClient
}

func InitOrderClient() OrderClient {

	orderConn, err := grpc.Dial(config.GetString("rpc_order"), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	orderGrpc := pb.NewOrderServiceClient(orderConn)

	return OrderClient{OrderService: orderGrpc}
}
