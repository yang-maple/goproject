package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStudentLogic {
	return &AddStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------student-----------------------
func (l *AddStudentLogic) AddStudent(in *pb.AddStudentReq) (*pb.AddStudentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddStudentResp{}, nil
}
