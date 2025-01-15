package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchClusterInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchClusterInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchClusterInfoLogic {
	return &SearchClusterInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchClusterInfoLogic) SearchClusterInfo(in *pb.SearchClusterInfoReq) (*pb.SearchClusterInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchClusterInfoResp{}, nil
}
