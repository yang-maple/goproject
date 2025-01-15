package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAddrsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAddrsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAddrsLogic {
	return &UpdateAddrsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAddrsLogic) UpdateAddrs(in *pb.UpdateAddrsReq) (*pb.UpdateAddrsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateAddrsResp{}, nil
}
