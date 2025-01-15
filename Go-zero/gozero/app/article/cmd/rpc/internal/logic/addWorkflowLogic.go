package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddWorkflowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddWorkflowLogic {
	return &AddWorkflowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------workflow-----------------------
func (l *AddWorkflowLogic) AddWorkflow(in *pb.AddWorkflowReq) (*pb.AddWorkflowResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddWorkflowResp{}, nil
}
