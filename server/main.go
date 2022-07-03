package main

import (
	"github.com/RyotaKITA-12/hack-a-matcher.git/routes"
    "github.com/RyotaKITA-12/hack-a-matcher.git/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()
	app.Use(cors.New(cors.Config{AllowCredentials: true}))

	routes.Setup(app)

	app.Listen(":8888")
}
