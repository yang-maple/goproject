package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStudentCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStudentCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStudentCourseLogic {
	return &UpdateStudentCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStudentCourseLogic) UpdateStudentCourse(in *pb.UpdateStudentCourseReq) (*pb.UpdateStudentCourseResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateStudentCourseResp{}, nil
}
