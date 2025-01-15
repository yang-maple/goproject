package logic

import (
	"context"

	"gozero/app/article/cmd/rpc/internal/svc"
	"gozero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchAlertsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchAlertsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchAlertsLogic {
	return &SearchAlertsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchAlertsLogic) SearchAlerts(in *pb.SearchAlertsReq) (*pb.SearchAlertsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchAlertsResp{}, nil
}
