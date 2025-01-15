package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCourseByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCourseByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCourseByIdLogic {
	return &GetCourseByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCourseByIdLogic) GetCourseById(in *pb.GetCourseByIdReq) (*pb.GetCourseByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetCourseByIdResp{}, nil
}
