package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAlertsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAlertsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAlertsLogic {
	return &AddAlertsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------alerts-----------------------
func (l *AddAlertsLogic) AddAlerts(in *pb.AddAlertsReq) (*pb.AddAlertsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddAlertsResp{}, nil
}
