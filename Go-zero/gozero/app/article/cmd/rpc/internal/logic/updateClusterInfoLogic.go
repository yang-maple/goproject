package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateClusterInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateClusterInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateClusterInfoLogic {
	return &UpdateClusterInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateClusterInfoLogic) UpdateClusterInfo(in *pb.UpdateClusterInfoReq) (*pb.UpdateClusterInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateClusterInfoResp{}, nil
}
