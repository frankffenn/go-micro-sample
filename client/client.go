// Code generated by rpc-gen. DO NOT EDIT.

package main

import (
	context "context"
	mod "github.com/frankffenn/go-micro-sample/mod"
	client "github.com/micro/go-micro/v2/client"
)

// Client API for HelloServer service

type HelloServerService struct {
	c    client.Client
	name string
}

func NewHelloServerService(name string, c client.Client) *HelloServerService {
	return &HelloServerService{
		c:    c,
		name: name,
	}
}

func requestOpts() []client.RequestOption {
	return []client.RequestOption{
		client.WithContentType("application/json"),
	}
}

func (c *HelloServerService) Say(ctx context.Context, in *mod.SayRequest, opts ...client.CallOption) (*mod.SayResponse, error) {
	req := c.c.NewRequest(c.name, "HelloServer.Say", in, requestOpts()...)
	out := new(mod.SayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *HelloServerService) Run(ctx context.Context, in *mod.RunRequest, opts ...client.CallOption) (*mod.RunResponse, error) {
	req := c.c.NewRequest(c.name, "HelloServer.Run", in, requestOpts()...)
	out := new(mod.RunResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}