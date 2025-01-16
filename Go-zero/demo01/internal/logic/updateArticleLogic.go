package logic

import (
	"context"

	"demo01/internal/svc"
	"demo01/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateArticleLogic) UpdateArticle(in *pb.UpdateArticleReq) (*pb.UpdateArticleResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateArticleResp{}, nil
}
