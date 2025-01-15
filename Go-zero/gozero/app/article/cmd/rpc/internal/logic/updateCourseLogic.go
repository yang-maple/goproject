package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCourseLogic {
	return &UpdateCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCourseLogic) UpdateCourse(in *pb.UpdateCourseReq) (*pb.UpdateCourseResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateCourseResp{}, nil
}
