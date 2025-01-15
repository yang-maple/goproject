package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Params struct {
	Width, Height int
}

type MyService interface {
	Area(Params, *int) error
	Perimeter(Params, *int) error
}

func Register(server MyService) {
	rpc.Register(server)
}

type NewClient struct {
	c *rpc.Client
}

func (c *NewClient) Area(p Params, res *int) error {
	return c.c.Call("Rect.Area", p, &res)
}

func (c *NewClient) Perimeter(p Params, res *int) error {
	return c.c.Call("Rect.Perimeter", p, &res)
}
func MyClient(addr string) (*NewClient, error) {
	c, err := jsonrpc.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &NewClient{c}, nil
}
