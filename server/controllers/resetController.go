package controllers

import (
	"fmt"
	"log"
	"math/rand"
	"net/smtp"

	"github.com/RyotaKITA-12/hack-a-matcher.git/database"
	"github.com/RyotaKITA-12/hack-a-matcher.git/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SendMail(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	token := ReturnRandomString(12)
	resetPass := models.ResetPassword{
		Email: data["email"],
		Token: token,
	}
	database.DB.Create(&resetPass)

	from := "hack.a.matcher@gmail.com"
	to := []string{data["email"]}

	sendFrom := fmt.Sprintf("From: %s\n", from)
	subject := fmt.Sprintf("Subject: %s\n", "Password reset.")
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	url := "http://localhost:8080/reset/" + token
	message := fmt.Sprintf("Click <a href=\"%s\">here</a> to reset password.", url)
	err := smtp.SendMail(
		"smtp:1025",
		nil,
		from,
		to,
		[]byte(sendFrom+subject+mime+message),
	)

	if err != nil {
		log.Fatalln(err)
	}

	return c.JSON(fiber.Map{"message": "Success."})
}

func Reset(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "Password do not match."})
	}

	var resetPass = models.ResetPassword{}

	err := database.DB.Where("token=?", data["token"]).Last(&resetPass)
	if err.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "Inbalid token"})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	database.DB.Model(&models.User{}).Where("email=?", resetPass.Email).Update("password", password)

	return c.JSON(fiber.Map{"message": "Success."})
}

func ReturnRandomString(n int) string {
	var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randStr := make([]rune, n)
	for i := range randStr {
		randStr[i] = runes[rand.Intn(len(runes))]
	}

	return string(randStr)
}
