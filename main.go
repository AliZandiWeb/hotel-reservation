package main

import (
	"flag"

	"github.com/AliZandiWeb/hotel-reservation/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "the listen address of the api server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("api/v1")

	apiv1.Get("/user", api.HandlerGetUsers)
	apiv1.Get("/user/:id", api.HandlerGetUserById)

	app.Listen(*listenAddr)
}
