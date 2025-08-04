package main

import (
	"github.com/gofiber/fiber/v2"
	"NoteApp-microservice/user_service/handlers"
	"NoteApp-microservice/user_service/database"
	"log"
)


func main() {
	database.Migrate()

	app := fiber.New()
    
	app.Post("/register",handlers.RegisterUser)
	app.Post("/login",handlers.Login)
    // Fiber'in kendi server'ını kullan
    log.Fatal(app.Listen(":8080"))
}