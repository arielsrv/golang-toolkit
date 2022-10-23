package main

import (
	"github.com/arielsrv/golang-toolkit/webserver"
	"log"
	"net/http"
)

func main() {
	application := &webserver.Application{
		UseRecovery:  true,
		UseRequestID: true,
		UseLogger:    true,
		UseSwagger:   true,
	}

	application.Register(http.MethodGet, "/ping", func(ctx *webserver.Context) error {
		return ctx.SendString("pong")
	})

	log.Fatal(application.Start("localhost:8080"))
}
