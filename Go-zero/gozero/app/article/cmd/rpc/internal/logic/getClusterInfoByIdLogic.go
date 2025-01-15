package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClusterInfoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetClusterInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClusterInfoByIdLogic {
	return &GetClusterInfoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetClusterInfoByIdLogic) GetClusterInfoById(in *pb.GetClusterInfoByIdReq) (*pb.GetClusterInfoByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetClusterInfoByIdResp{}, nil
}
