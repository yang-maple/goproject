package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateWorkflowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateWorkflowLogic {
	return &UpdateWorkflowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateWorkflowLogic) UpdateWorkflow(in *pb.UpdateWorkflowReq) (*pb.UpdateWorkflowResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateWorkflowResp{}, nil
}
