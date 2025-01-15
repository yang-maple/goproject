package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentCourseByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStudentCourseByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentCourseByIdLogic {
	return &GetStudentCourseByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStudentCourseByIdLogic) GetStudentCourseById(in *pb.GetStudentCourseByIdReq) (*pb.GetStudentCourseByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetStudentCourseByIdResp{}, nil
}
