package api

import (
	"github.com/Nasa28/hotel-reservation-api/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "Chinasa",
		LastName:  "Kalu",
	}
	return c.JSON(user)
}

func HandleGetUserById(c *fiber.Ctx) error {
	return c.JSON("James")
}
