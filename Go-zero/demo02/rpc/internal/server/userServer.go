// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.5
// Source: user.proto

package server

import (
	"context"

	"demo02/rpc/internal/logic"
	"demo02/rpc/internal/svc"
	"demo02/rpc/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

// -----------------------user-----------------------
func (s *UserServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) AddUser(ctx context.Context, in *pb.AddUserReq) (*pb.AddUserResp, error) {
	l := logic.NewAddUserLogic(ctx, s.svcCtx)
	return l.AddUser(in)
}

func (s *UserServer) UpdateUser(ctx context.Context, in *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
	l := logic.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}

func (s *UserServer) DelUser(ctx context.Context, in *pb.DelUserReq) (*pb.DelUserResp, error) {
	l := logic.NewDelUserLogic(ctx, s.svcCtx)
	return l.DelUser(in)
}

func (s *UserServer) GetUserById(ctx context.Context, in *pb.GetUserByIdReq) (*pb.GetUserByIdResp, error) {
	l := logic.NewGetUserByIdLogic(ctx, s.svcCtx)
	return l.GetUserById(in)
}

func (s *UserServer) SearchUser(ctx context.Context, in *pb.SearchUserReq) (*pb.SearchUserResp, error) {
	l := logic.NewSearchUserLogic(ctx, s.svcCtx)
	return l.SearchUser(in)
}