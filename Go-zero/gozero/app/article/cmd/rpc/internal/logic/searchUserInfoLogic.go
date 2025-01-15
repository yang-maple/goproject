package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserInfoLogic {
	return &SearchUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserInfoLogic) SearchUserInfo(in *pb.SearchUserInfoReq) (*pb.SearchUserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchUserInfoResp{}, nil
}
