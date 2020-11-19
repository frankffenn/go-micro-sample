package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/micro/go-micro/v2/web"
	"time"
)

func main() {

	service := web.NewService(
		web.Name("go.micro.srv.gate"),
		web.RegisterTTL(30 *time.Second),
		web.RegisterInterval(20 *time.Second),
		web.Address(":8888"),
	)

	service.Init()

	c := cors.New(cors.Options{
		AllowedOrigins:         []string{"*"},
		AllowedMethods:         []string{"GET","POST","OPTIONS"},
		AllowedHeaders:         []string{"*"},
		MaxAge:                 3600,
		Debug:                  false,
	})
	r := mux.NewRouter()
	service.Handle("/", c.Handler(r))
	// register routers

	// Run the server
	service.Run()
}