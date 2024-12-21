package api

import (
	"github.com/AliZandiWeb/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	store *db.Store
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

func (h *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID()
	}
	filter := bson.M{"hotelID": oid}
	rooms, err := h.store.Room.GetRooms(c.Context(), filter)
	if err != nil {
		return err
	}
	return c.JSON(rooms)
}

func (h *HotelHandler) HandleGetHotelByID(c *fiber.Ctx) error {
	id := c.Params("id")
	// oid, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return ErrInvalidID()
	// }
	hotel, err := h.store.Hotel.GetHotelByID(c.Context(), id)
	if err != nil {
		return ErrResourceNotFound("hotel")
	}
	return c.JSON(hotel)
}

// pagination -> hotel_store.go & db.go

type ResourceResponse struct {
	Data    any `json:"data"`
	Results int `json:""results"`
	Page    int `json:"page"`
}
type HotelQueryParams struct {
	db.Pagination
	Rating int
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {

	var params HotelQueryParams
	if err := c.QueryParser(&params); err != nil {
		return ErrBadRequest()
	}
	// filter
	filter := db.Map{
		"rating": params.Rating,
	}
	hotels, err := h.store.Hotel.GetHotels(c.Context(), filter, &params.Pagination)
	if err != nil {
		return ErrResourceNotFound("hotels")
	}
	resp := ResourceResponse{
		Results: len(hotels),
		Data:    hotels,
		Page:    int(params.Page),
	}
	return c.JSON(resp)
}
