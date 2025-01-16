package logic

import (
	"context"

	"demo02/models"
	"demo02/rpc_gorm/internal/svc"
	"demo02/rpc_gorm/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserLogic {
	return &DelUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserLogic) DelUser(in *pb.DelUserReq) (*pb.DelUserResp, error) {
	// todo: add your logic here and delete this line
	user := models.User{
		Id: in.Id,
	}
	result := l.svcCtx.DB.Take(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	l.svcCtx.DB.Delete(&models.User{}, in.Id)
	return &pb.DelUserResp{}, nil
}