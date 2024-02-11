package main

import (
	"flag"
	"log"

	"github.com/Nasa28/hotel-reservation-api/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	port := flag.String("port", ":4000", "The listen address of the server")

	app := fiber.New()
	flag.Parse()

	apiV1 := app.Group("api/v1")

	apiV1.Get("/user", api.HandleGetUsers)
	apiV1.Get("/user:id", api.HandleGetUserById)
	err := app.Listen(*port)
	if err != nil {
		log.Fatalf("Error while starting server: %v", err)
	}
}
