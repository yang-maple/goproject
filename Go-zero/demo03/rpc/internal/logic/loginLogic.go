package logic

import (
	"context"
	"errors"

	"demo03/models"
	"demo03/rpc/internal/svc"
	"demo03/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// --login--
func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// todo: add your logic here and delete this line
	user := models.User{
		Username: in.Username,
	}
	tx := l.svcCtx.DB.Take(&user)
	if tx.Error != nil {
		return nil, errors.New("XUserNotExist")
	}
	if user.Password != in.Password {
		return nil, errors.New("XPasswordError")
	}
	return &pb.LoginResp{}, nil
}
