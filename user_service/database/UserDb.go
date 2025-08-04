package database

import (
	"os"
	"log"
	"NoteApp-microservice/user_service/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	err error

	U models.User
)

func Migrate() {
	dsn := os.Getenv("DB_DNS")
	if dsn == "" {
		log.Fatal("DB_DNS environment variable is not set")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&U)
}

func AddUser(user *models.User) error {
    return db.Create(user).Error
}

func DeleteUser() {
	db.Delete(&U,U.ID)
}

func GetUser(where ...interface{}) models.User {
	var user models.User
	db.First(&user, where...)
	return user
}

func GetAllUsers(where ...interface{}) []models.User {
	var users []models.User
	db.Find(&users,where...)
	return users
}