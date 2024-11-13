package api

import (
	"github.com/AliZandiWeb/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func HandlerGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Ali",
		LastName:  "Zandi",
	}
	return c.JSON(u)
}
func HandlerGetUserById(c *fiber.Ctx) error {
	return c.JSON("Ali")
}
