package services

import (
	"context"
	"fmt"

	"github.com/INVITATION-RPC-GATEWAY/api/app/auth/client"
	"github.com/INVITATION-RPC-GATEWAY/api/app/auth/pb"
)

type AuthService interface {
	Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error)
	Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error)
	Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error)
	RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error)
}

type authService struct {
	authClient client.AuthClient
}

func NewAuthService(authClient client.AuthClient) AuthService {
	return &authService{authClient: authClient}
}

func (m *authService) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	out, err := m.authClient.AuthService.Register(ctx, in)
	if err != nil {
		fmt.Printf("err register %s", err.Error())
		return out, err
	}

	return out, err
}

func (m *authService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {

	out, err := m.authClient.AuthService.Login(ctx, in)
	if err != nil {
		fmt.Printf("err login %s", err.Error())
		return out, err
	}

	return out, err
}

func (m *authService) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error) {

	out, err := m.authClient.AuthService.Logout(ctx, in)
	if err != nil {
		fmt.Printf("err logout %s", err.Error())
		return out, err
	}

	return out, err
}

func (m *authService) RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {

	out, err := m.authClient.AuthService.RefreshToken(ctx, in)
	if err != nil {
		fmt.Printf("err refresh token %s", err.Error())
		return out, err
	}

	return out, err
}
