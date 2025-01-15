package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAlertsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAlertsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAlertsLogic {
	return &UpdateAlertsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAlertsLogic) UpdateAlerts(in *pb.UpdateAlertsReq) (*pb.UpdateAlertsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateAlertsResp{}, nil
}
