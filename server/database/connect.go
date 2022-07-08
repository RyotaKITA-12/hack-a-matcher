package database

import (
	"fmt"
	"log"
	"os"

	"github.com/RyotaKITA-12/hack-a-matcher.git/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	schema         = "%s:%s@tcp(mysql:3306)/%s?charset=utf8&parseTime=True&loc=Local"
	username       = os.Getenv("MYSQL_USER")
	password       = os.Getenv("MYSQL_PASSWORD")
	dbName         = os.Getenv("MYSQL_DATABASE")
	datasourceName = fmt.Sprintf(schema, username, password, dbName)

	DB *gorm.DB
)

func Connect() {
	db, err := gorm.Open(mysql.Open(datasourceName), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	DB = db

	db.AutoMigrate(&models.User{}, &models.ResetPassword{}, &models.Profile{}, &models.Link{}, &models.Skill{}, &models.Role{})
}
