package main

import (
	"context"
	"flag"
	"log"

	"github.com/AliZandiWeb/hotel-reservation/api"
	"github.com/AliZandiWeb/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userColl = "users"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "the listen address of the api server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	// handler intialization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))
	app := fiber.New(config)
	apiv1 := app.Group("api/v1")

	apiv1.Get("/user", userHandler.HandlerGetUsers)
	apiv1.Get("/user/:id", userHandler.HandlerGetUserByID)
	apiv1.Post("/user", userHandler.HandlerPostUser)

	app.Listen(*listenAddr)
}
