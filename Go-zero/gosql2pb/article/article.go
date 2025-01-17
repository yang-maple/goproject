// Code generated by goctl. DO NOT EDIT.
// Source: article.proto

package article

import (
	"context"

	"gosql2pb/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCusterReq      = pb.AddCusterReq
	AddCusterResp     = pb.AddCusterResp
	Custer            = pb.Custer
	DelCusterReq      = pb.DelCusterReq
	DelCusterResp     = pb.DelCusterResp
	GetCusterByIdReq  = pb.GetCusterByIdReq
	GetCusterByIdResp = pb.GetCusterByIdResp
	SearchCusterReq   = pb.SearchCusterReq
	SearchCusterResp  = pb.SearchCusterResp
	UpdateCusterReq   = pb.UpdateCusterReq
	UpdateCusterResp  = pb.UpdateCusterResp

	Article interface {
		// -----------------------Custer-----------------------
		AddCuster(ctx context.Context, in *AddCusterReq, opts ...grpc.CallOption) (*AddCusterResp, error)
		UpdateCuster(ctx context.Context, in *UpdateCusterReq, opts ...grpc.CallOption) (*UpdateCusterResp, error)
		DelCuster(ctx context.Context, in *DelCusterReq, opts ...grpc.CallOption) (*DelCusterResp, error)
		GetCusterById(ctx context.Context, in *GetCusterByIdReq, opts ...grpc.CallOption) (*GetCusterByIdResp, error)
		SearchCuster(ctx context.Context, in *SearchCusterReq, opts ...grpc.CallOption) (*SearchCusterResp, error)
	}

	defaultArticle struct {
		cli zrpc.Client
	}
)

func NewArticle(cli zrpc.Client) Article {
	return &defaultArticle{
		cli: cli,
	}
}

// -----------------------Custer-----------------------
func (m *defaultArticle) AddCuster(ctx context.Context, in *AddCusterReq, opts ...grpc.CallOption) (*AddCusterResp, error) {
	client := pb.NewArticleClient(m.cli.Conn())
	return client.AddCuster(ctx, in, opts...)
}

func (m *defaultArticle) UpdateCuster(ctx context.Context, in *UpdateCusterReq, opts ...grpc.CallOption) (*UpdateCusterResp, error) {
	client := pb.NewArticleClient(m.cli.Conn())
	return client.UpdateCuster(ctx, in, opts...)
}

func (m *defaultArticle) DelCuster(ctx context.Context, in *DelCusterReq, opts ...grpc.CallOption) (*DelCusterResp, error) {
	client := pb.NewArticleClient(m.cli.Conn())
	return client.DelCuster(ctx, in, opts...)
}

func (m *defaultArticle) GetCusterById(ctx context.Context, in *GetCusterByIdReq, opts ...grpc.CallOption) (*GetCusterByIdResp, error) {
	client := pb.NewArticleClient(m.cli.Conn())
	return client.GetCusterById(ctx, in, opts...)
}

func (m *defaultArticle) SearchCuster(ctx context.Context, in *SearchCusterReq, opts ...grpc.CallOption) (*SearchCusterResp, error) {
	client := pb.NewArticleClient(m.cli.Conn())
	return client.SearchCuster(ctx, in, opts...)
}
