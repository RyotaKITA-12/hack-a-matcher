package routes

import (
	"github.com/RyotaKITA-12/hack-a-matcher.git/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/api/user", controllers.User)
	app.Get("/api/profile", controllers.Profile)
	app.Get("/api/link", controllers.Link)
	app.Get("/api/post", controllers.Post)
	app.Get("/api/username", controllers.UserName)
	app.Get("/api/skill", controllers.Skill)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/password/forgot", controllers.SendMail)
	app.Post("/api/password/reset", controllers.Reset)
	app.Post("/api/logout", controllers.Logout)
	app.Post("/api/register/profile", controllers.RegisterProfile)
	app.Post("/api/register/link", controllers.RegisterLink)
	app.Post("/api/register/post", controllers.RegisterPost)
	app.Post("/api/register/post/skill", controllers.RegisterSkill)
}
