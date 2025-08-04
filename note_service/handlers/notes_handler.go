package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"NoteApp-microservice/note_service/database"
	"NoteApp-microservice/note_service/data"
	"NoteApp-microservice/note_service/models"
	"github.com/golang-jwt/jwt"
)

func getJWTFromHeaderOrCookie(c *fiber.Ctx) string {
	authHeader := c.Get("Authorization")
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		return authHeader[7:]
	}
	return c.Cookies("jwt")
}

func GetNote(c *fiber.Ctx) error {
	tokenString := c.Cookies("jwt")
    if tokenString == "" {
        return c.Redirect("/login")
    }
	id := c.Params("id")
	note1 := database.GetNote("id = ?",id)
	return c.JSON(note1)
}

func GetAllNotes(c *fiber.Ctx) error {
	tokenString := c.Cookies("jwt")
    if tokenString == "" {
        return c.Redirect("/login")
    }
	notes := database.GetAllNotes()
	return c.JSON(notes)
}

func AddNote(c *fiber.Ctx) error {
	tokenString := getJWTFromHeaderOrCookie(c)
    if tokenString == "" {
        return c.Redirect("/login")
    }

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fiber.ErrUnauthorized
        }
        return []byte("super-secret-key"), nil
    })

    if err != nil || !token.Valid {
        return c.Redirect("/login")
    }

    // JWT claims'lerini al
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return c.Redirect("/login")
    }

    // userId claim'inden kullanıcıyı al
    userID := int(claims["userId"].(float64))

	req := new(data.NoteData)
	if err := c.BodyParser(req); err != nil {
			fmt.Println("Error parsing request body:", err)
			return c.Status(400).SendString("Invalid request format")
	}

	note1 := &models.Note{
		UserID: uint(userID),
		Title: req.Title,
		Content: req.Content,
	}

	err1 := database.AddNote(note1)
	if err1!= nil {
		fmt.Println("Failed Add Note to Database")
	}

	return c.Status(fiber.StatusOK).SendString("Succesfull Add Note Database")
}

func UpdateNotes(c *fiber.Ctx) error {
	tokenString := c.Cookies("jwt")
    if tokenString == "" {
        return c.Redirect("/login")
    }

	id := c.Params("id")
	note1 := database.GetNote("id = ?",id)
	req := new(data.NoteData)

	if err := c.BodyParser(req); err != nil {
		fmt.Println("Error parsing request body",err)
		return c.Status(400).SendString("Invalid request format")
	}

	note1.Content = req.Content
	note1.Title = req.Title

	database.UpdateNotes(note1)
	return c.Status(fiber.StatusOK).SendString("Succesfull Uptade Note")
}

func DeleteNote(c *fiber.Ctx) error {
	tokenString := c.Cookies("jwt")
    if tokenString == "" {
        return c.Redirect("/login")
    }
	
	id := c.Params("id")
	note1 := database.GetNote(id)
	if note1.ID == 0 {
        return c.Status(404).SendString("Note not found")
    }
	database.N = note1
	
	database.DeleteNote()
	
	return c.Status(fiber.StatusOK).SendString("Succesfull Delete")
}
