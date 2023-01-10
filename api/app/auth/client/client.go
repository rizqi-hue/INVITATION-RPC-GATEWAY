package client

import (
	"github.com/INVITATION-RPC-GATEWAY/api/app/auth/pb"
	"github.com/INVITATION-RPC-GATEWAY/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	AuthService pb.AuthServiceClient
}

func InitAuthClient() AuthClient {
	// ---- start dialing auth grpc ----
	authConn, err := grpc.Dial(config.GetString("rpc_auth"), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	authGrpc := pb.NewAuthServiceClient(authConn)

	return AuthClient{AuthService: authGrpc}
}
