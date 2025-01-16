package logic

import (
	"context"

	"demo01/internal/svc"
	"demo01/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchArticleLogic {
	return &SearchArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchArticleLogic) SearchArticle(in *pb.SearchArticleReq) (*pb.SearchArticleResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchArticleResp{}, nil
}