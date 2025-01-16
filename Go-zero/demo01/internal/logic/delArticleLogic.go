package logic

import (
	"context"

	"demo01/internal/svc"
	"demo01/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelArticleLogic {
	return &DelArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelArticleLogic) DelArticle(in *pb.DelArticleReq) (*pb.DelArticleResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelArticleResp{
		Message: "success",
		Code:    200,
		Data:    "del",
	}, nil
}
