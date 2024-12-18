package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AliZandiWeb/hotel-reservation/api"
	"github.com/AliZandiWeb/hotel-reservation/db"
	"github.com/AliZandiWeb/hotel-reservation/db/fixtures"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	// for drop hotel with same name ("after make seed he make double hotel and room and this code tone make same hotel and room")
	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
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

}
