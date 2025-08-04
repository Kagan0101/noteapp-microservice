package handlers

import (
	"NoteApp-microservice/user_service/auth"
	"NoteApp-microservice/user_service/database"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"NoteApp-microservice/user_service/models"
	"github.com/gofiber/fiber/v2"
	"NoteApp-microservice/user_service/utils"
)


func RegisterUser(c *fiber.Ctx) error {
    req := new(auth.RegisterRequest)
    if err := c.BodyParser(req); err != nil {
        return c.Status(400).SendString("Invalid request format")
    }
	hash := sha256.Sum256([]byte(req.Password))
	hexed := hex.EncodeToString(hash[:])
	fmt.Println(hexed)
	newUser := &models.User{
		Username: req.Username,
		Email: req.Email,
		Password: hexed,
	}
	err := database.AddUser(newUser)
	if err != nil{
		fmt.Println("Failed add user to database")
	}
    return c.Status(fiber.StatusOK).SendString("User registered") 
}

func Login(c *fiber.Ctx) error {
	req := new(auth.LoginRequest)
	if err := c.BodyParser(req); err != nil {
        return c.Status(400).SendString("Invalid request format")
    }
	hash := sha256.Sum256([]byte(req.Password))
	hexed := hex.EncodeToString(hash[:])
	user := database.GetUser("email = ? AND password = ?",req.Email,hexed)
	if user.ID == 0 {
		return c.Status(404).SendString("User or Password Wrong")
	}else{
		token, err := utils.GenerateToken(&user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
        c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			HTTPOnly: !c.IsFromLocal(),
			Secure:   !c.IsFromLocal(),
			MaxAge:   3600 * 24 * 7, // 7 days
		})
        return c.JSON(fiber.Map{
			"token" : token,
		})
	}
}


