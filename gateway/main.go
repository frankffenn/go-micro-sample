package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
	"github.com/rs/cors"
	"log"
	"time"
)

func main() {

	service := web.NewService(
		web.Name("go.micro.srv.gateway"),
		web.RegisterTTL(30*time.Second),
		web.RegisterInterval(20*time.Second),
		web.Address(":7777"),
	)

	service.Init()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		MaxAge:         3600,
		Debug:          false,
	})

	r := gin.New()
	RegisterHandler(r)
	service.Handle("/", c.Handler(r))

	// Run the server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
