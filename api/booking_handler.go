package api

import (
	// "net/http"

	"github.com/AliZandiWeb/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type BookingHandler struct {
	store *db.Store
}

func NewBookingHandler(store *db.Store) *BookingHandler {
	return &BookingHandler{
		store: store,
	}
}
func (h *BookingHandler) HandleCancelBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.store.Booking.GetBookingByID(c.Context(), id)
	if err != nil {
		return ErrResourceNotFound("booking")
	}
	user, err := getAuthUser(c)
	if err != nil {
		return ErrUnAuthorized()
	}

	if booking.UserID != user.ID {
		return ErrUnAuthorized()
		// return c.Status(http.StatusUnauthorized).JSON(genericResp{
		// 	Type: "error",
		// 	Msg:  "not authorized",
		// })
	}
	if err := h.store.Booking.UpdateBooking(c.Context(), c.Params("id"), bson.M{"canceled": true}); err != nil {
		return err
	}

	return c.JSON(genericResp{
		Type: "msg",
		Msg:  "updated",
	})
}

// TODO : this needs to be admin authorized!
func (h *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	bookings, err := h.store.Booking.GetBookings(c.Context(), bson.M{})
	if err != nil {
		return ErrResourceNotFound("booking")
	}
	return c.JSON(bookings)
}

// TODO : this needs to be user authorized!
func (h *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.store.Booking.GetBookingByID(c.Context(), id)
	if err != nil {
		return ErrResourceNotFound("booking")
	}

	user, err := getAuthUser(c)
	if err != nil {
		return ErrUnAuthorized()
	}
	if booking.UserID != user.ID {
		return ErrUnAuthorized()
		// return c.Status(http.StatusUnauthorized).JSON(genericResp{
		// 	Type: "error",
		// 	Msg:  "unauthorized",
		// })
	}
	return c.JSON(booking)
}
