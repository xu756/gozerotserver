package logic

import (
	"context"

	"gozerotserver/app/user/api/internal/svc"
	"gozerotserver/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.Res, err error) {

	return &types.Res{
		Code: 200,
		Msg:  "success",
		Data: l.ctx.Value("user"),
	}, nil
}