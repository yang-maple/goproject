package logic

import (
	"context"

	"demo02/rpc_db/internal/svc"
	"demo02/rpc_db/pb"

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
	err := l.svcCtx.UserModel.Delete(context.TODO(), in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DelUserResp{}, nil
}
