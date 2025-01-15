package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelWorkflowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelWorkflowLogic {
	return &DelWorkflowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelWorkflowLogic) DelWorkflow(in *pb.DelWorkflowReq) (*pb.DelWorkflowResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelWorkflowResp{}, nil
}
