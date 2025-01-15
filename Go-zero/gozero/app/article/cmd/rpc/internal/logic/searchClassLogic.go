package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchClassLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchClassLogic {
	return &SearchClassLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchClassLogic) SearchClass(in *pb.SearchClassReq) (*pb.SearchClassResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchClassResp{}, nil
}
