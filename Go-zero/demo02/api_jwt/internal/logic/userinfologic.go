package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"demo02/api_jwt/internal/svc"
	"demo02/api_jwt/internal/types"

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
	userId := l.ctx.Value("user_id").(json.Number)
	fmt.Printf("%v, %T, \n", userId, userId)
	username := l.ctx.Value("username").(string)
	uid, _ := userId.Int64()

	return &types.UserInfoResponse{
		Id:       uint(uid),
		UserName: username,
	}, nil
}
