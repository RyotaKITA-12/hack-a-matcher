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
	view_time, _ := time.Parse("2006-01-02", data["view_period"])

	post := models.Post{
		UserID:     user_id,
		UserName:   data["user_name"],
		Post_type:  data["post_type"],
		Title:      data["title"],
		Content:    data["content"],
		ViewPeriod: view_time,
	}
	database.DB.Create(&post)

	return c.JSON(post)
}

func Skill(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("post_id"))
	if err != nil {
		log.Println(err)
	}

	var skills []models.PostSkill
	database.DB.Where("post_id=?", id).Find(&skills)

	return c.JSON(skills)
}

func RegisterSkill(c *fiber.Ctx) error {
	var data map[string][]string

	if err := c.BodyParser(&data); err != nil {
		log.Fatalln(err)
	}

	var post models.Post
	database.DB.Where("user_id = ?", data["user_id"][0]).Last(&post)
	post_id := int(post.ID)

	for _, skill := range data["skills"] {
		log.Println(skill)
		skill := models.PostSkill{
			PostID: post_id,
			Lang:   skill,
		}
		database.DB.Create(&skill)
	}

	return c.JSON(post)
}
