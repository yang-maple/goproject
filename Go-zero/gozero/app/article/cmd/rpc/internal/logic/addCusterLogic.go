package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCusterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCusterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCusterLogic {
	return &AddCusterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------Custer-----------------------
func (l *AddCusterLogic) AddCuster(in *pb.AddCusterReq) (*pb.AddCusterResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddCusterResp{}, nil
}
