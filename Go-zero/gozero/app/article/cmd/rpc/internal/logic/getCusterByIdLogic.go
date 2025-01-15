package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCusterByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCusterByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCusterByIdLogic {
	return &GetCusterByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCusterByIdLogic) GetCusterById(in *pb.GetCusterByIdReq) (*pb.GetCusterByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetCusterByIdResp{}, nil
}