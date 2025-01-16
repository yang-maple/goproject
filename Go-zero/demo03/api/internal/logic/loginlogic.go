package logic

import (
	"context"

	"demo03/api/internal/svc"
	"demo03/api/internal/types"
	"demo03/common/jwts"
	"demo03/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.UserRpc.Login(context.TODO(), &user.LoginReq{
		Username: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return "", err
	}
	auth := l.svcCtx.Config.Auth
	token, err := jwts.CreateToken(jwts.JwtPayLoad{
		UserID:   1,
		Username: req.UserName,
		Role:     1,
	}, auth.AccessSecret, auth.AccessExpire)
	if err != nil {
		return "", err
	}
	return token, nil
}
