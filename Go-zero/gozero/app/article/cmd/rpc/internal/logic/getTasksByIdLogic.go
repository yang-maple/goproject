package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTasksByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTasksByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTasksByIdLogic {
	return &GetTasksByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTasksByIdLogic) GetTasksById(in *pb.GetTasksByIdReq) (*pb.GetTasksByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetTasksByIdResp{}, nil
}
