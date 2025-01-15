package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchStudentLogic {
	return &SearchStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchStudentLogic) SearchStudent(in *pb.SearchStudentReq) (*pb.SearchStudentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchStudentResp{}, nil
}
