package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelClusterInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelClusterInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelClusterInfoLogic {
	return &DelClusterInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelClusterInfoLogic) DelClusterInfo(in *pb.DelClusterInfoReq) (*pb.DelClusterInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelClusterInfoResp{}, nil
}
