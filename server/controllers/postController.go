package controllers

import (
	"log"
	"strconv"
	"time"

	"github.com/RyotaKITA-12/hack-a-matcher.git/database"
	"github.com/RyotaKITA-12/hack-a-matcher.git/models"
	"github.com/gofiber/fiber/v2"
)

func Post(c *fiber.Ctx) error {
	now_time := time.Now()

	var posts []models.Post
	database.DB.Where("view_period > ?", now_time).Find(&posts)

    // data, _ := json.Marshal(posts)
    // log.Printf("%s", data)

	return c.JSON(posts)
}

func RegisterPost(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		log.Fatalln(err)
	}

	user_id, _ := strconv.Atoi(data["user_id"])
	rec_num, _ := strconv.Atoi(data["recruit_num"])
	view_time, _ := time.Parse("2006-01-02", data["view_period"])

	post := models.Post{
		UserID:     user_id,
		Title:      data["title"],
		RecruitNum: rec_num,
		Content:    data["content"],
		ViewPeriod: view_time,
	}
	database.DB.Create(&post)

	return c.JSON(post)
}
