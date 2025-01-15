package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddClassLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddClassLogic {
	return &AddClassLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------class-----------------------
func (l *AddClassLogic) AddClass(in *pb.AddClassReq) (*pb.AddClassResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddClassResp{}, nil
}
