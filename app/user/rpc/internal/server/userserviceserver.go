// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"gozerotserver/app/user/rpc/internal/logic"
	"gozerotserver/app/user/rpc/internal/svc"
	"gozerotserver/app/user/rpc/user"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceServer) Login(ctx context.Context, in *user.User) (*user.UserResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}
