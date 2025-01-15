package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelClassLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelClassLogic {
	return &DelClassLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelClassLogic) DelClass(in *pb.DelClassReq) (*pb.DelClassResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelClassResp{}, nil
}
