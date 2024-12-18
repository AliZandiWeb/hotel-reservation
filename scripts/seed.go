package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AliZandiWeb/hotel-reservation/api"
	"github.com/AliZandiWeb/hotel-reservation/db"
	"github.com/AliZandiWeb/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client       *mongo.Client
	roomStore    db.RoomStore
	hotelStore   db.HotelStore
	ctx          = context.Background()
	userStore    db.UserStore
	bookingStore db.BookingStore
)

func seedUser(isAdmin bool, fname, lname, email, password string) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		Email:     email,
		FirstName: fname,
		LastName:  lname,
		Password:  password,
	})
	if err != nil {
		log.Fatal(err)
	}
	user.IsAdmin = isAdmin
	insertedUser, err := userStore.InsertUser(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s -> %s\n", user.Email, api.CreateTokenFromUser(user))
	return insertedUser
}
func seedHotel(name, location string, rating int) *types.Hotel {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}
	return insertedHotel
}
func seedRoom(size string, ss bool, price float64, hotelID primitive.ObjectID) *types.Room {
	room := &types.Room{
		Size:    size,
		SeaSide: ss,
		Price:   price,
		HotelID: hotelID,
	}
	insertedRoom, err := roomStore.InsertRoom(context.Background(), room)
	if err != nil {
		log.Fatal(err)
	}
	return insertedRoom
}
func seedBooking(userrID, roomID primitive.ObjectID, numPersons int, from, till time.Time) *types.Booking {
	booking := &types.Booking{
		UserID:     userrID,
		RoomID:     roomID,
		NumPersons: numPersons,
		FromDate:   from,
		TillDate:   till,
	}
	insertedBooking, err := bookingStore.InsertBooking(context.Background(), booking)
	if err != nil {
		log.Fatal(err)
	}
	return insertedBooking
}
func main() {
	ali := seedUser(false, "ali", "zandi", "ali@gmail.com", "supersecurepassword")
	seedUser(true, "admin", "admin", "admin@admin.com", "adminpassword")
	seedHotel("Eram", "Iran", 3)
	hotel := seedHotel("Bellucia", "France", 5)
	seedHotel("borochenka", "Italy", 4)
	seedRoom("small", false, 8.99, hotel.ID)
	room := seedRoom("medium", false, 18.99, hotel.ID)
	seedRoom("large", true, 38.99, hotel.ID)
	seedBooking(ali.ID, room.ID, 4, time.Now(), time.Now().AddDate(0, 0, 2))
}

func init() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	// for drop hotel with same name ("after make seed he make double hotel and room and this code tone make same hotel and room")
	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
	userStore = db.NewMongoUserStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
	bookingStore = db.NewMongoBookingStore(client)
}
