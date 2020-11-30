package main

import (
	"github.com/frankffenn/go-micro-sample/hello-srv/impl"
	"github.com/frankffenn/go-micro-sample/rpc-gen"
)

func main() {
	if err := rpc_gen.WriteServicesToFile("../../client/client.go", "main",
		impl.HelloServer{},
	); err != nil {
		panic(err)
	}
}
