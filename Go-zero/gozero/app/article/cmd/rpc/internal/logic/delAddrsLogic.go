package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAddrsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAddrsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAddrsLogic {
	return &DelAddrsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAddrsLogic) DelAddrs(in *pb.DelAddrsReq) (*pb.DelAddrsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelAddrsResp{}, nil
}
