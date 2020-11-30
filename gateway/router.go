package main

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type AppJWTMiddleware struct {
	*jwt.GinJWTMiddleware
}

var (
	TokenTimeout        = time.Hour * 24
	TokenRefreshTimeout = time.Hour * 24 * 30
	AuthUserMiddleware  *AppJWTMiddleware
)

func RegisterHandler(r *gin.Engine) {
	apiV1 := r.Group("/api/v1")

	AuthUserMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm: "User",
		Key:   []byte("fake news"),
	})

	if err != nil {
		log.Fatal(err)
	}

	apiV1.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, ResponseSuccess(&jsonObj{"pong": true}))
	})

	// checking token below
	r.Use(AuthUserMiddleware.MiddlewareFunc())
}
