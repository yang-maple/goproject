package logic

import (
	"context"

	"demo03/models"
	"demo03/rpc/internal/svc"
	"demo03/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------user-----------------------
func (l *AddUserLogic) AddUser(in *pb.AddUserReq) (*pb.AddUserResp, error) {
	// todo: add your logic here and delete this line
	user := models.User{
		Username: in.Username,
		Password: in.Password,
	}
	err := l.svcCtx.DB.Create(&user)
	if err.Error != nil {
		return nil, err.Error
	}
	return &pb.AddUserResp{}, nil
}
