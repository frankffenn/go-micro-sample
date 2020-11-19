package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/frankffenn/go-micro-sample/mod"
)

func main() {

	service := micro.NewService()
	service.Init()

	client := NewHelloServerService("HelloServer", service.Client())

	rsp, err := client.Say(context.Background(), &mod.SayRequest{Name:"frank"})
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