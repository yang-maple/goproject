package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCourseLogic {
	return &AddCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------course-----------------------
func (l *AddCourseLogic) AddCourse(in *pb.AddCourseReq) (*pb.AddCourseResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddCourseResp{}, nil
}
