package controllers

import (
	"log"
	"strconv"
	"time"

	"github.com/RyotaKITA-12/hack-a-matcher.git/database"
	"github.com/RyotaKITA-12/hack-a-matcher.git/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Profile(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	log.Println(cookie)
	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Un autenticated."})
	}

	claims := token.Claims.(*Claims)
	id := claims.Issuer

	var profile models.Profile
	database.DB.Where("user_id=?", id).First(&profile)

	return c.JSON(profile)
}

func RegisterProfile(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		log.Fatalln(err)
	}

	i, _ := strconv.Atoi(data["user_id"])
	t, _ := time.Parse("20060102", data["birthday"])

	profile := models.Profile{
		UserID:     i,
		Email:      data["email"],
		Birth:      t,
		Address:    data["address"],
		Occupation: data["occupation"],
	}
	database.DB.Create(&profile)

	return c.JSON(profile)
}

func Link(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	log.Println(cookie)
	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Un autenticated."})
	}

	claims := token.Claims.(*Claims)
	id := claims.Issuer

	var link models.Link
	database.DB.Where("user_id=?", id).First(&link)

	return c.JSON(link)
}

func RegisterLink(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		log.Fatalln(err)
	}

	i, _ := strconv.Atoi(data["user_id"])

	link := models.Link{
		UserID:    i,
		Portfolio: data["portfolio"],
		Twitter:   data["twitter"],
		Github:    data["github"],
	}

	database.DB.Create(&link)

	return c.JSON(link)
}
