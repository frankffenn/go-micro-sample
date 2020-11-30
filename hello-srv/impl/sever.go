package impl

import (
	"context"
	"fmt"

	"github.com/frankffenn/go-micro-sample/mod"
)

type HelloServer struct{}

func NewServer() *HelloServer {
	return &HelloServer{}
}

func (s *HelloServer) Say(ctx context.Context, req *mod.SayRequest, rsp *mod.SayResponse) error {
	rsp.Reply = fmt.Sprintf("fake news!%s", req.Name)
	return nil
}

func (s *HelloServer) Run(ctx context.Context, req *mod.RunRequest, rsp *mod.RunResponse) error {
	rsp.Something = "the young won’t speak wood"
	return nil
}
