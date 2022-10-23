package main

import (
	"github.com/arielsrv/golang-toolkit/webserver/api"
	"log"
	"net/http"
)

func main() {
	application := &api.Application{
		UseRecovery:  true,
		UseRequestID: true,
		UseLogger:    true,
		UseSwagger:   true,
	}

	application.Register(http.MethodGet, "/ping", func(ctx *api.Context) error {
		return ctx.SendString("pong")
	})

	log.Fatal(application.Start("localhost:8080"))
}
