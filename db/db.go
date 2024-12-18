package db

const (
	DBNAME     = "hotel-reservation"
	TestDBNAME = "hotel-reservation-test"
	DBURI      = "mongodb://localhost:27017"
)

type Store struct {
	User    UserStore
	Hotel   HotelStore
	Room    RoomStore
	Booking BookingStore
}

// func ToObjectID(id string) (primitive.ObjectID, error) {
// 	oid, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return oid
// }
