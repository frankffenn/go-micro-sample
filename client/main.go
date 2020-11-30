package main

import (
	"context"
	"fmt"
	"github.com/frankffenn/go-micro-sample/mod"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
)

func main() {
	register := etcdv3.NewRegistry()
	service := micro.NewService(
		micro.Registry(register),
	)
	service.Init()

	client := NewHelloServerService("go.micro.srv.hello", service.Client())

	rsp, err := client.Say(context.Background(), &mod.SayRequest{Name: "frank"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.Reply)

	rsp1, err := client.Run(context.Background(), &mod.RunRequest{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp1.Something)
}
