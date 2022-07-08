package controllers

import (
	"log"
	"strconv"
	"time"

	"github.com/RyotaKITA-12/hack-a-matcher.git/database"
	"github.com/RyotaKITA-12/hack-a-matcher.git/models"
	"github.com/gofiber/fiber/v2"
)

func RegisterProfile(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		log.Fatalln(err)
	}

	i, _ := strconv.Atoi(data["user_id"])
	t, _ := time.Parse("20060102", data["birthday"])
	// i, _ := strconv.Atoi("3")
	// t, _ := time.Parse("20060102", "20011212")

	profile := models.Profile{
		UserID:     i,
		Email:      data["email"],
		Birth:      t,
		Address:    data["address"],
		Occupation: data["occupation"],
	}
	// profile := models.Profile{
	// 	UserID:     i,
	// 	Email:      "test@example.com",
	// 	Birth:      t,
	// 	Address:    "東京都",
	// 	Occupation: "学生",
	// }

	database.DB.Create(&profile)

	return c.JSON(profile)
}
