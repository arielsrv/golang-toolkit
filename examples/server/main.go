package main

import (
	"github.com/arielsrv/golang-toolkit/server"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

func main() {
	app := server.New()

	handler := func(ctx *fiber.Ctx) error {
		return ctx.SendString("pong")
	}

	server.RegisterHandler(handler)

	app.Add(http.MethodGet, "/ping", server.Use(handler))

	log.Fatal(app.Start("localhost:8080"))
}
