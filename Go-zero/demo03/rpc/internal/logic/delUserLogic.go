package logic

import (
	"context"

	"demo03/models"
	"demo03/rpc/internal/svc"
	"demo03/rpc/pb"

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
	err := l.svcCtx.DB.Where("id = ?", in.Id).Delete(&models.User{})
	if err.Error != nil {
		return nil, err.Error
	}
	return &pb.DelUserResp{}, nil
}
