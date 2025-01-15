package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleByIdLogic {
	return &GetArticleByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetArticleByIdLogic) GetArticleById(in *pb.GetArticleByIdReq) (*pb.GetArticleByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetArticleByIdResp{}, nil
}
