package article

import (
	"context"

	"gozero/app/article/cmd/api/internal/svc"
	"gozero/app/article/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteArticleLogic) DeleteArticle(req *types.DeleteArticleReq) (resp *types.DeleteArticleResp, err error) {
	// todo: add your logic here and delete this line

	return
}
