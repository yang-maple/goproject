package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchStudentCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchStudentCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchStudentCourseLogic {
	return &SearchStudentCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchStudentCourseLogic) SearchStudentCourse(in *pb.SearchStudentCourseReq) (*pb.SearchStudentCourseResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchStudentCourseResp{}, nil
}
