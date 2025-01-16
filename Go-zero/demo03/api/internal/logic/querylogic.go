package logic

import (
	"context"

	"demo03/api/internal/svc"
	"demo03/api/internal/types"
	"demo03/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLogic {
	return &QueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryLogic) Query(req *types.QueryRequest) (resp string, err error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.UserRpc.GetUserById(context.TODO(), &user.GetUserByIdReq{
		Id: int64(req.Id),
	})
	if err != nil {
		return "", err
	}
	return response.String(), nil
}
