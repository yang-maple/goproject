// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/hellomsdn.proto

package hellomsdn

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Hellomsdn service

func NewHellomsdnEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Hellomsdn service

type HellomsdnService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Hellomsdn_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Hellomsdn_PingPongService, error)
}

type hellomsdnService struct {
	c    client.Client
	name string
}

func NewHellomsdnService(name string, c client.Client) HellomsdnService {
	return &hellomsdnService{
		c:    c,
		name: name,
	}
}

func (c *hellomsdnService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Hellomsdn.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hellomsdnService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Hellomsdn_StreamService, error) {
	req := c.c.NewRequest(c.name, "Hellomsdn.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &hellomsdnServiceStream{stream}, nil
}

type Hellomsdn_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type hellomsdnServiceStream struct {
	stream client.Stream
}

func (x *hellomsdnServiceStream) Close() error {
	return x.stream.Close()
}

func (x *hellomsdnServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *hellomsdnServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *hellomsdnServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *hellomsdnServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hellomsdnService) PingPong(ctx context.Context, opts ...client.CallOption) (Hellomsdn_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Hellomsdn.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &hellomsdnServicePingPong{stream}, nil
}

type Hellomsdn_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type hellomsdnServicePingPong struct {
	stream client.Stream
}

func (x *hellomsdnServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *hellomsdnServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *hellomsdnServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *hellomsdnServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *hellomsdnServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *hellomsdnServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Hellomsdn service

type HellomsdnHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Hellomsdn_StreamStream) error
	PingPong(context.Context, Hellomsdn_PingPongStream) error
}

func RegisterHellomsdnHandler(s server.Server, hdlr HellomsdnHandler, opts ...server.HandlerOption) error {
	type hellomsdn interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Hellomsdn struct {
		hellomsdn
	}
	h := &hellomsdnHandler{hdlr}
	return s.Handle(s.NewHandler(&Hellomsdn{h}, opts...))
}

type hellomsdnHandler struct {
	HellomsdnHandler
}

func (h *hellomsdnHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.HellomsdnHandler.Call(ctx, in, out)
}

func (h *hellomsdnHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.HellomsdnHandler.Stream(ctx, m, &hellomsdnStreamStream{stream})
}

type Hellomsdn_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type hellomsdnStreamStream struct {
	stream server.Stream
}

func (x *hellomsdnStreamStream) Close() error {
	return x.stream.Close()
}

func (x *hellomsdnStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *hellomsdnStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *hellomsdnStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *hellomsdnStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *hellomsdnHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.HellomsdnHandler.PingPong(ctx, &hellomsdnPingPongStream{stream})
}

type Hellomsdn_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type hellomsdnPingPongStream struct {
	stream server.Stream
}

func (x *hellomsdnPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *hellomsdnPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *hellomsdnPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *hellomsdnPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *hellomsdnPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *hellomsdnPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
