package main

import (
	"context"
	"log"
	"os"

	"github.com/AliZandiWeb/hotel-reservation/api"
	"github.com/AliZandiWeb/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// configuration
// 1. MongoDB endpoint
// 2. ListenAddress of our HTTP server
// 3. JWT secret
// 4. MongoDBName

var config = fiber.Config{
	ErrorHandler: api.Errorhandler,
}

func main() {
	// 2024-12-14 19:12:34.041709145 +0330 +0330 m=+0.001915540
	// listenAddr := flag.String("listenAddr", ":5000", "the listen address of the api server")
	// flag.Parse()
	mongoEndpoint := os.Getenv("MONGO_DB_URL")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoEndpoint))
	if err != nil {
		log.Fatal(err)
	}

	// handler intialization
	var (
		hotelStore   = db.NewMongoHotelStore(client)
		roomStore    = db.NewMongoRoomStore(client, hotelStore)
		userStore    = db.NewMongoUserStore(client)
		bookingStore = db.NewMongoBookingStore(client)
		store        = &db.Store{
			Hotel:   hotelStore,
			Room:    roomStore,
			User:    userStore,
			Booking: bookingStore,
		}
		userHandler    = api.NewUserHandler(userStore)
		hotelHandler   = api.NewHotelHandler(store)
		authHandler    = api.NewAuthHandler(userStore)
		roomHandler    = api.NewRoomHandler(store)
		bookingHandler = api.NewBookingHandler(store)
		app            = fiber.New(config)
		auth           = app.Group("/api")
		apiv1          = app.Group("/api/v1", api.JWTAuthentication(userStore))
		admin          = apiv1.Group("/admin", api.AdminAuth)
	)
	// auth
	auth.Post("/auth", authHandler.HandleAuthenticate)
	// Versioned API routes
	// user Handlers
	apiv1.Get("/user", userHandler.HandlerGetUsers)
	apiv1.Get("/user/:id", userHandler.HandlerGetUserByID)
	apiv1.Post("/user", userHandler.HandlerPostUser)
	apiv1.Delete("/user/:id", userHandler.HandlerDeleteUser)
	apiv1.Put("/user/:id", userHandler.HandlerPutUser)
	// hotel Handlers
	apiv1.Get("/hotel", hotelHandler.HandleGetHotels)
	apiv1.Get("/hotel/:id", hotelHandler.HandleGetHotelByID)
	apiv1.Get("/hotel/:id/rooms", hotelHandler.HandleGetRooms)
	// room Handlers
	apiv1.Get("/room", roomHandler.HandleGetRooms)
	apiv1.Post("/room/:id/book", roomHandler.HandleBookRoom)

	// // Booking Handlers
	apiv1.Get("/booking/:id", bookingHandler.HandleGetBooking)
	apiv1.Get("/booking/:id/cancel", bookingHandler.HandleCancelBooking)
	//  Admin
	admin.Get("/booking", bookingHandler.HandleGetBookings)

	listenAddr := os.Getenv("HTTP_LISTEN_ADDRESS")
	app.Listen(listenAddr)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
