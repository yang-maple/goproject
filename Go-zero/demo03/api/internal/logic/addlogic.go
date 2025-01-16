package logic

import (
	"context"

	"demo03/api/internal/svc"
	"demo03/api/internal/types"
	"demo03/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddRequest) (resp string, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.UserRpc.AddUser(context.TODO(), &user.AddUserReq{
		Username: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return "", err
	}
	return "success", nil
}
