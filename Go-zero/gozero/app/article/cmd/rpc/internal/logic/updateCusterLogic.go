package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCusterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCusterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCusterLogic {
	return &UpdateCusterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCusterLogic) UpdateCuster(in *pb.UpdateCusterReq) (*pb.UpdateCusterResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateCusterResp{}, nil
}
