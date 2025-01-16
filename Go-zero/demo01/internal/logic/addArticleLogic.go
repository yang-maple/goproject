package logic

import (
	"context"

	"demo01/internal/svc"
	"demo01/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleLogic {
	return &AddArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------article-----------------------
func (l *AddArticleLogic) AddArticle(in *pb.AddArticleReq) (*pb.AddArticleResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddArticleResp{
		Message: "success",
		Code:    200,
		Data:    "add",
	}, nil
}
