package main

import (
	"github.com/frankffenn/go-micro-sample/hello-srv/impl"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
)

func main() {
	register := etcdv3.NewRegistry()
	service := micro.NewService(
		micro.Registry(register),
		micro.Name("go.micro.srv.hello"),
	)

	service.Init()

	micro.RegisterHandler(service.Server(), impl.NewServer())

	// Run the server
	service.Run()
}
