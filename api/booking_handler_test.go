package api

import (
	"fmt"
	"testing"
	"time"

	"github.com/AliZandiWeb/hotel-reservation/db/fixtures"
)

func TestAdminGetBooking(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)

	user := fixtures.AddUser(db.Store, "ali", "zandi", false)
	hotel := fixtures.AddHotel(db.Store, "bar hotel ", "a", 4, nil)
	room := fixtures.AddRoom(db.Store, "small", true, 4.4, hotel.ID)

	from := time.Now()
	till := time.Now().AddDate(0, 0, 4)

	booking := fixtures.AddBooking(db.Store, user.ID, room.ID, 2, from, till)

	fmt.Println(booking)
}
