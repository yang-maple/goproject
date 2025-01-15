package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAddrsByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAddrsByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAddrsByIdLogic {
	return &GetAddrsByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAddrsByIdLogic) GetAddrsById(in *pb.GetAddrsByIdReq) (*pb.GetAddrsByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetAddrsByIdResp{}, nil
}
