package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelStudentLogic {
	return &DelStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelStudentLogic) DelStudent(in *pb.DelStudentReq) (*pb.DelStudentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelStudentResp{}, nil
}
