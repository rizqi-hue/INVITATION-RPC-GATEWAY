package client

import (
	"github.com/INVITATION-RPC-GATEWAY/api/app/product/pb"
	"github.com/INVITATION-RPC-GATEWAY/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductClient struct {
	ProductService pb.ProductServiceClient
}

func InitProductClient() ProductClient {
	// ---- start dialing product grpc ----
	productConn, err := grpc.Dial(config.GetString("rpc_product"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	productGrpc := pb.NewProductServiceClient(productConn)

	return ProductClient{ProductService: productGrpc}
}
