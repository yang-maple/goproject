package logic

import (
	"context"
	"fmt"

	"demo02/models"
	"demo02/rpc_db/internal/svc"
	"demo02/rpc_db/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------user-----------------------
func (l *AddUserLogic) AddUser(in *pb.AddUserReq) (*pb.AddUserResp, error) {
	// todo: add your logic here and delete this line
	result, error := l.svcCtx.UserModel.Insert(context.TODO(), &models.User{
		Username: in.Username,
		Password: in.Password,
	})
	fmt.Println(result.RowsAffected())
	if error != nil {
		return nil, error
	}
	return &pb.AddUserResp{}, nil
}
