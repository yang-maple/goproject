package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTasksLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTasksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTasksLogic {
	return &AddTasksLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------tasks-----------------------
func (l *AddTasksLogic) AddTasks(in *pb.AddTasksReq) (*pb.AddTasksResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddTasksResp{}, nil
}
