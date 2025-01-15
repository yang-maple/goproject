package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStudentCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddStudentCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStudentCourseLogic {
	return &AddStudentCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------studentCourse-----------------------
func (l *AddStudentCourseLogic) AddStudentCourse(in *pb.AddStudentCourseReq) (*pb.AddStudentCourseResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddStudentCourseResp{}, nil
}
