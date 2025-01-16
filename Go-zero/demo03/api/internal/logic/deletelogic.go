package logic

import (
	"context"

	"demo03/api/internal/svc"
	"demo03/api/internal/types"
	"demo03/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.DeleteRequest) (resp string, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.UserRpc.DelUser(context.TODO(), &user.DelUserReq{
		Id: int64(req.Id),
	})
	if err != nil {
		return "", err
	}
	return "success", nil
}
