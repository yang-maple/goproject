package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCourseLogic {
	return &DelCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCourseLogic) DelCourse(in *pb.DelCourseReq) (*pb.DelCourseResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelCourseResp{}, nil
}
