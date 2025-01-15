package logic

import (
	"context"

	"gosql2pb/internal/svc"
	"gosql2pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCusterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchCusterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCusterLogic {
	return &SearchCusterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchCusterLogic) SearchCuster(in *pb.SearchCusterReq) (*pb.SearchCusterResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchCusterResp{}, nil
}
