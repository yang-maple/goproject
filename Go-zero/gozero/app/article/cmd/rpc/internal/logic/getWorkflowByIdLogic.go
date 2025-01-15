package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWorkflowByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWorkflowByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWorkflowByIdLogic {
	return &GetWorkflowByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetWorkflowByIdLogic) GetWorkflowById(in *pb.GetWorkflowByIdReq) (*pb.GetWorkflowByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetWorkflowByIdResp{}, nil
}
