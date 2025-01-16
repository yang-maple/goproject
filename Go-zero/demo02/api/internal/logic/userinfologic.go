package logic

import (
	"context"

	"demo02/api/internal/svc"
	"demo02/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return &types.UserInfoResponse{
		Code: 200,
		Msg:  "success",
		Data: types.UserInfo{
			Addr:     "111",
			UserName: "222",
		},
	}, nil
}
