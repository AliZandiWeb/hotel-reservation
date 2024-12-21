package db

// const (
//
//	DBNAME     = "hotel-reservation"
//	TestDBNAME = "hotel-reservation-test"
//	DBURI      = "mongodb://localhost:27017"
//
// )
const MongoDBNameEnvName = "MONGO_DB_NAME"

// pagination
type Pagination struct {
	Page  int64
	Limit int64
}

type Store struct {
	User    UserStore
	Hotel   HotelStore
	Room    RoomStore
	Booking BookingStore
}
