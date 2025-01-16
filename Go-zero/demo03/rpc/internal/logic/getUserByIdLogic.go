package logic

import (
	"context"

	"demo03/models"
	"demo03/rpc/internal/svc"
	"demo03/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *pb.GetUserByIdReq) (*pb.GetUserByIdResp, error) {
	// todo: add your logic here and delete this line
	user := models.User{}
	err := l.svcCtx.DB.Where("id = ?", in.Id).Take(&user)
	if err.Error != nil {
		return nil, err.Error
	}
	return &pb.GetUserByIdResp{
		User: &pb.User{
			Username: user.Username,
			Password: user.Password,
		},
	}, nil
}
