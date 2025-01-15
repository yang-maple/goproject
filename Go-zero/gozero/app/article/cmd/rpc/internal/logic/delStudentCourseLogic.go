package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelStudentCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelStudentCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelStudentCourseLogic {
	return &DelStudentCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelStudentCourseLogic) DelStudentCourse(in *pb.DelStudentCourseReq) (*pb.DelStudentCourseResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelStudentCourseResp{}, nil
}
