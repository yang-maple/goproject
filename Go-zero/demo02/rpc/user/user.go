// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.5
// Source: user.proto

package user

import (
	"context"

	"demo02/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddUserReq      = pb.AddUserReq
	AddUserResp     = pb.AddUserResp
	DelUserReq      = pb.DelUserReq
	DelUserResp     = pb.DelUserResp
	GetUserByIdReq  = pb.GetUserByIdReq
	GetUserByIdResp = pb.GetUserByIdResp
	LoginReq        = pb.LoginReq
	LoginResp       = pb.LoginResp
	SearchUserReq   = pb.SearchUserReq
	SearchUserResp  = pb.SearchUserResp
	UpdateUserReq   = pb.UpdateUserReq
	UpdateUserResp  = pb.UpdateUserResp
	User            = pb.User

	UserZrpcClient interface {
		// -----------------------user-----------------------
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserResp, error)
		UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error)
		DelUser(ctx context.Context, in *DelUserReq, opts ...grpc.CallOption) (*DelUserResp, error)
		GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error)
		SearchUser(ctx context.Context, in *SearchUserReq, opts ...grpc.CallOption) (*SearchUserResp, error)
	}

	defaultUserZrpcClient struct {
		cli zrpc.Client
	}
)

func NewUserZrpcClient(cli zrpc.Client) UserZrpcClient {
	return &defaultUserZrpcClient{
		cli: cli,
	}
}

// -----------------------user-----------------------
func (m *defaultUserZrpcClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUserZrpcClient) AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.AddUser(ctx, in, opts...)
}

func (m *defaultUserZrpcClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

func (m *defaultUserZrpcClient) DelUser(ctx context.Context, in *DelUserReq, opts ...grpc.CallOption) (*DelUserResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.DelUser(ctx, in, opts...)
}

func (m *defaultUserZrpcClient) GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GetUserById(ctx, in, opts...)
}

func (m *defaultUserZrpcClient) SearchUser(ctx context.Context, in *SearchUserReq, opts ...grpc.CallOption) (*SearchUserResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.SearchUser(ctx, in, opts...)
}
