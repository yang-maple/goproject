package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchAddrsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchAddrsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchAddrsLogic {
	return &SearchAddrsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchAddrsLogic) SearchAddrs(in *pb.SearchAddrsReq) (*pb.SearchAddrsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchAddrsResp{}, nil
}
