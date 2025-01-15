// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/helloword.proto

package helloword

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

// Api Endpoints for Helloword service

func NewHellowordEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Helloword service

type HellowordService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Helloword_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Helloword_PingPongService, error)
}

type hellowordService struct {
	c    client.Client
	name string
}

func NewHellowordService(name string, c client.Client) HellowordService {
	return &hellowordService{
		c:    c,
		name: name,
	}
}

func (c *hellowordService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Helloword.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hellowordService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Helloword_StreamService, error) {
	req := c.c.NewRequest(c.name, "Helloword.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &hellowordServiceStream{stream}, nil
}

type Helloword_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type hellowordServiceStream struct {
	stream client.Stream
}

func (x *hellowordServiceStream) Close() error {
	return x.stream.Close()
}

func (x *hellowordServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *hellowordServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *hellowordServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *hellowordServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hellowordService) PingPong(ctx context.Context, opts ...client.CallOption) (Helloword_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Helloword.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &hellowordServicePingPong{stream}, nil
}

type Helloword_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type hellowordServicePingPong struct {
	stream client.Stream
}

func (x *hellowordServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *hellowordServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *hellowordServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *hellowordServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *hellowordServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *hellowordServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Helloword service

type HellowordHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Helloword_StreamStream) error
	PingPong(context.Context, Helloword_PingPongStream) error
}

func RegisterHellowordHandler(s server.Server, hdlr HellowordHandler, opts ...server.HandlerOption) error {
	type helloword interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Helloword struct {
		helloword
	}
	h := &hellowordHandler{hdlr}
	return s.Handle(s.NewHandler(&Helloword{h}, opts...))
}

type hellowordHandler struct {
	HellowordHandler
}

func (h *hellowordHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.HellowordHandler.Call(ctx, in, out)
}

func (h *hellowordHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.HellowordHandler.Stream(ctx, m, &hellowordStreamStream{stream})
}

type Helloword_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type hellowordStreamStream struct {
	stream server.Stream
}

func (x *hellowordStreamStream) Close() error {
	return x.stream.Close()
}

func (x *hellowordStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *hellowordStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *hellowordStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *hellowordStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *hellowordHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.HellowordHandler.PingPong(ctx, &hellowordPingPongStream{stream})
}

type Helloword_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type hellowordPingPongStream struct {
	stream server.Stream
}

func (x *hellowordPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *hellowordPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *hellowordPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *hellowordPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *hellowordPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *hellowordPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
