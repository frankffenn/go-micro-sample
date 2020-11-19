package main

import (
	"github.com/frankffenn/go-micro-sample/rpc-gen"
	"github.com/frankffenn/go-micro-sample/hello-srv"
)

func main () {
	if err := rpc_gen.WriteServicesToFile("../../client/client.go", "main",
		hello.HelloServer{},
	); err != nil {
		panic(err)
	}
}