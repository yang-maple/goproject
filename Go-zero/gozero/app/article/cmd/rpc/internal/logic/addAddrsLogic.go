package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAddrsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAddrsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAddrsLogic {
	return &AddAddrsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------addrs-----------------------
func (l *AddAddrsLogic) AddAddrs(in *pb.AddAddrsReq) (*pb.AddAddrsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddAddrsResp{}, nil
}
