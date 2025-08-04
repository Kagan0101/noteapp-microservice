package main

import (
	"github.com/gofiber/fiber/v2"
	"NoteApp-microservice/note_service/handlers"
	"NoteApp-microservice/note_service/database"
	"log"
)


func main() {
	database.Migrate()

	app := fiber.New()
    
	app.Get("/notes/:id",handlers.GetNote)
	app.Get("/",handlers.GetAllNotes)
	app.Post("/notes/add",handlers.AddNote)
	app.Post("/notes/uptade/:id",handlers.UpdateNotes)
	app.Delete("/notes/delete/:id",handlers.DeleteNote)
    // Fiber'in kendi server'ını kullan
    log.Fatal(app.Listen(":8081"))
}