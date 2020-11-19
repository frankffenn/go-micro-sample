package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/frankffenn/go-micro-sample/hello-srv"
)

func main() {

	service := micro.NewService(
		micro.Name("HelloServer"),
	)

	service.Init()

	micro.RegisterHandler(service.Server(), hello.NewServer())


	// Run the server
	service.Run()
}