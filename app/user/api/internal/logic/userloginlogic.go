package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"gozerotserver/app/user/api/internal/svc"
	"gozerotserver/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *UserLoginLogic) getJwtToken(secretKey string, iat int64, seconds int64, username string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["user"] = username
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
func (l *UserLoginLogic) UserLogin(req *types.UserLogin) (resp *types.Res, err error) {
	now := time.Now().Unix()
	jwtToken, e := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, req.Username)
	if e != nil {
		return nil, err
	}
	return &types.Res{
		Code: 200,
		Msg:  "success",
		Data: jwtToken,
	}, nil
}
