package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserInfoLogic {
	return &DelUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserInfoLogic) DelUserInfo(in *pb.DelUserInfoReq) (*pb.DelUserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelUserInfoResp{}, nil
}
