package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateClassLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateClassLogic {
	return &UpdateClassLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateClassLogic) UpdateClass(in *pb.UpdateClassReq) (*pb.UpdateClassResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateClassResp{}, nil
}
