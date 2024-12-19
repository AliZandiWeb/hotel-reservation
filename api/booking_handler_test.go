package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/AliZandiWeb/hotel-reservation/api/middleware"
	"github.com/AliZandiWeb/hotel-reservation/db/fixtures"
	"github.com/AliZandiWeb/hotel-reservation/types"
)

func TestUserGetBooking(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)
	var (
		noneAuthUser = fixtures.AddUser(db.Store, "jimmi", "watercooler", false)
		user         = fixtures.AddUser(db.Store, "ali", "zandi", false)
		hotel        = fixtures.AddHotel(db.Store, "bar hotel ", "a", 4, nil)
		room         = fixtures.AddRoom(db.Store, "small", true, 4.4, hotel.ID)

		from = time.Now()
		till = time.Now().AddDate(0, 0, 4)

		booking        = fixtures.AddBooking(db.Store, user.ID, room.ID, 2, from, till)
		app            = fiber.New()
		route          = app.Group("/:id", middleware.JWTAuthentication(db.User))
		bookingHandler = NewBookingHandler(db.Store)
	)
	route.Get("/", bookingHandler.HandleGetBooking)
	// fmt.Println(booking.ID.String())
	req := httptest.NewRequest("GET", fmt.Sprintf("/%s", booking.ID.Hex()), nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(user))
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("none 200 response got %d", resp.StatusCode)
	}
	var bookingResp *types.Booking
	if err := json.NewDecoder(resp.Body).Decode(&bookingResp); err != nil {
		t.Fatal(err)
	}
	// fmt.Println(bookingResp)
	if bookingResp.ID != booking.ID {
		t.Fatalf("expected %s got %s", booking.ID, bookingResp.ID)
	}
	if bookingResp.UserID != booking.UserID {
		t.Fatalf("expected %s got %s", booking.UserID, bookingResp.UserID)
	}
	// ////////////////////////////////////////////////////////////////
	req = httptest.NewRequest("GET", fmt.Sprintf("/%s", booking.ID.Hex()), nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(noneAuthUser))
	resp, err = app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("none 200 response got %d", resp.StatusCode)
	}
}

func TestAdminGetBooking(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)
	var (
		adminUser = fixtures.AddUser(db.Store, "admin", "admin", true)
		user      = fixtures.AddUser(db.Store, "ali", "zandi", false)
		hotel     = fixtures.AddHotel(db.Store, "bar hotel ", "a", 4, nil)
		room      = fixtures.AddRoom(db.Store, "small", true, 4.4, hotel.ID)

		from = time.Now()
		till = time.Now().AddDate(0, 0, 4)

		booking        = fixtures.AddBooking(db.Store, user.ID, room.ID, 2, from, till)
		app            = fiber.New()
		admin          = app.Group("/", middleware.JWTAuthentication(db.User), middleware.AdminAuth)
		bookingHandler = NewBookingHandler(db.Store)
	)
	// fmt.Println(booking)

	_ = booking

	admin.Get("/", bookingHandler.HandleGetBookings)
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(adminUser))
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("none 200 response got %d", resp.StatusCode)
	}
	var bookings []*types.Booking
	if err := json.NewDecoder(resp.Body).Decode(&bookings); err != nil {
		t.Fatal(err)
	}
	if len(bookings) != 1 {
		t.Fatalf("expected 1 booking got %d", len(bookings))
	}
	// if !reflect.DeepEqual(booking, bookings[0]) {
	// 	t.Fatalf("expected booking to be equal ")

	// }
	have := bookings[0]
	if have.ID != booking.ID {
		t.Fatalf("expected %s got %s", booking.ID, have.ID)
	}
	if have.UserID != booking.UserID {
		t.Fatalf("expected %s got %s", booking.UserID, have.UserID)
	}
	// test none-admin cannot access the bookings
	req = httptest.NewRequest("GET", "/", nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(user))
	resp, err = app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected a none 200 status code got %d", resp.StatusCode)
	}

}
