package api

import (
	"context"

	"github.com/AliZandiWeb/hotel-reservation/db"
	"github.com/AliZandiWeb/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandlerGetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := context.Background()
	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandlerGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Ali",
		LastName:  "Zandi",
	}
	return c.JSON(u)
}
