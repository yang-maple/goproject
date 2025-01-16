package logic

import (
	"context"

	"demo02/rpc_db/internal/svc"
	"demo02/rpc_db/pb"

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
	users, err := l.svcCtx.UserModel.FindOne(context.TODO(), in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserByIdResp{
		User: &pb.User{
			Id:       users.Id,
			Username: users.Username,
			Password: users.Password,
		},
	}, nil
}
