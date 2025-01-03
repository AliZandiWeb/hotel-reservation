package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/AliZandiWeb/hotel-reservation/api"
	"github.com/AliZandiWeb/hotel-reservation/db"
	"github.com/AliZandiWeb/hotel-reservation/db/fixtures"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	var (
		ctx           = context.Background()
		mongoEndpoint = os.Getenv("MONGO_DB_URL")
		mongoDBNAME   = os.Getenv("MONGO_DB_NAME")
	)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoEndpoint))
	if err != nil {
		log.Fatal(err)
	}
	// for drop hotel with same name ("after make seed he make double hotel and room and this code tone make same hotel and room")
	if err := client.Database(mongoDBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client)
	store := &db.Store{
		User:    db.NewMongoUserStore(client),
		Booking: db.NewMongoBookingStore(client),
		Room:    db.NewMongoRoomStore(client, hotelStore),
		Hotel:   db.NewMongoHotelStore(client),
	}
	user := fixtures.AddUser(store, "ali", "zandi", false)
	fmt.Println("ali -> ", api.CreateTokenFromUser(user))
	admin := fixtures.AddUser(store, "admin", "admin", true)
	fmt.Println("admin -> ", api.CreateTokenFromUser(admin))

	hotel := fixtures.AddHotel(store, "some hotel", "Bermuda", 3, nil)
	room := fixtures.AddRoom(store, "large", true, 123.56, hotel.ID)
	booking := fixtures.AddBooking(store, user.ID, room.ID, 3, time.Now(), time.Now().AddDate(0, 0, 3))
	fmt.Println(booking)

	for i := 0; i < 100; i++ {
		name := fmt.Sprintf("random hotel %d", i)
		location := fmt.Sprintf("location%d", i)
		fixtures.AddHotel(store, name, location, (rand.Intn(5) + 1), nil)
	}
}
