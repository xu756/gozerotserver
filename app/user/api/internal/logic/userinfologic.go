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

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfo, err error) {
	user, err := l.svcCtx.UserModel.FindOneByname(l.ctx, l.ctx.Value("user").(string))
	if err != nil {
		return nil, err
	}
	return &types.UserInfo{
		Username: user.Username,
		Password: user.Password,
	}, nil
}
