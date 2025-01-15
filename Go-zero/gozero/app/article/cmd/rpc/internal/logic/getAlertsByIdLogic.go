package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAlertsByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAlertsByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAlertsByIdLogic {
	return &GetAlertsByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAlertsByIdLogic) GetAlertsById(in *pb.GetAlertsByIdReq) (*pb.GetAlertsByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetAlertsByIdResp{}, nil
}
