package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCourseLogic {
	return &SearchCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchCourseLogic) SearchCourse(in *pb.SearchCourseReq) (*pb.SearchCourseResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchCourseResp{}, nil
}
