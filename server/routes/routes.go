package routes

import (
	"github.com/RyotaKITA-12/hack-a-matcher.git/controllers"
	"github.com/gofiber/fiber/v2"
)


func Setup(app *fiber.App) {
    app.Get("/api/user", controllers.User)
    app.Post("/api/register", controllers.Register)
    app.Post("/api/login", controllers.Login)
    app.Post("/api/password/forgot", controllers.SendMail)
    app.Post("/api/password/reset", controllers.Reset)
    app.Get("/api/logout", controllers.Logout)
}
