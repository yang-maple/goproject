package logic

import (
	"context"

	"demo02/models"
	"demo02/rpc_gorm/internal/svc"
	"demo02/rpc_gorm/pb"

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
	result := l.svcCtx.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &pb.AddUserResp{}, nil
}
