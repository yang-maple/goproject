package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchWorkflowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchWorkflowLogic {
	return &SearchWorkflowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchWorkflowLogic) SearchWorkflow(in *pb.SearchWorkflowReq) (*pb.SearchWorkflowResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchWorkflowResp{}, nil
}
