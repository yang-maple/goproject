package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetClassByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassByIdLogic {
	return &GetClassByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetClassByIdLogic) GetClassById(in *pb.GetClassByIdReq) (*pb.GetClassByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetClassByIdResp{}, nil
}
