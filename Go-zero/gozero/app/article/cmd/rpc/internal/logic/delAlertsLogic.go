package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAlertsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAlertsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAlertsLogic {
	return &DelAlertsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAlertsLogic) DelAlerts(in *pb.DelAlertsReq) (*pb.DelAlertsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelAlertsResp{}, nil
}
