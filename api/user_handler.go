package api

import (
	"errors"

	"github.com/AliZandiWeb/hotel-reservation/db"
	"github.com/AliZandiWeb/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandlerPutUser(c *fiber.Ctx) error {
	var (
		// values bson.M
		params types.UpdateUserParams
		userID = c.Params("id")
	)
	// oid, err := primitive.ObjectIDFromHex(userID)
	// if err != nil {
	// 	return err
	// }
	if err := c.BodyParser(&params); err != nil {
		return ErrBadRequest()
	}
	filter := db.Map{"_id": userID}
	if err := h.userStore.PutUser(c.Context(), filter, params); err != nil {
		return err
	}
	return c.JSON(map[string]string{"updated": userID})
}
func (h *UserHandler) HandlerDeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	if err := h.userStore.DeleteUser(c.Context(), userID); err != nil {
		return err
	}
	return c.JSON(map[string]string{"Deleted": userID})
}
func (h *UserHandler) HandlerPostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return ErrBadRequest()
	}
	if errors := params.Validate(); len(errors) > 0 {
		return c.JSON(errors)
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	insertedUser, err := h.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(insertedUser)
}
func (h *UserHandler) HandlerGetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{"error": "Not Found"})
		}
		return err
	}
	return c.JSON(user)
}
func (h *UserHandler) HandlerGetUsers(c *fiber.Ctx) error {
	user, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return ErrResourceNotFound("user")
	}
	return c.JSON(user)
}
