package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddClusterInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddClusterInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddClusterInfoLogic {
	return &AddClusterInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------clusterInfo-----------------------
func (l *AddClusterInfoLogic) AddClusterInfo(in *pb.AddClusterInfoReq) (*pb.AddClusterInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddClusterInfoResp{}, nil
}
