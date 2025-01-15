package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCusterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCusterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCusterLogic {
	return &DelCusterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCusterLogic) DelCuster(in *pb.DelCusterReq) (*pb.DelCusterResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelCusterResp{}, nil
}
