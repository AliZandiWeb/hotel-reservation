package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/AliZandiWeb/hotel-reservation/api"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(client)

	listenAddr := flag.String("listenAddr", ":5000", "the listen address of the api server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("api/v1")

	apiv1.Get("/user", api.HandlerGetUsers)
	apiv1.Get("/user/:id", api.HandlerGetUserById)

	app.Listen(*listenAddr)
}
