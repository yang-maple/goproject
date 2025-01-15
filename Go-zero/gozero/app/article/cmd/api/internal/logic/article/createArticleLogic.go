package article

import (
	"context"
	"gozero/app/article/cmd/rpc/article"

	"gozero/app/article/cmd/api/internal/svc"
	"gozero/app/article/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateArticleLogic {
	return &CreateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateArticleLogic) CreateArticle(req *types.CreateArticleReq) (resp *types.CreateArticleResp, err error) {
	// 这里就是调用rpc下的AddArticle方法
	_, err = l.svcCtx.ArticleRpc.AddArticle(l.ctx, &article.AddArticleReq{
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateArticleResp{}, nil
}
