// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userservice

import (
	"context"

	"gozerotserver/app/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	User         = user.User
	UserResponse = user.UserResponse

	UserService interface {
		Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

func (m *defaultUserService) Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}
