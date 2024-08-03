package routes

import (
    "notes/handlers"
    "github.com/gofiber/fiber/v2"
    "github.com/jmoiron/sqlx"
)

func SetupRoutes(app *fiber.App, db *sqlx.DB) {
    handler := handlers.NoteHandler{DB: db}

    app.Get("/notes", handler.GetNotes)
    app.Get("/notes/:id", handler.GetNoteByID)
    app.Post("/notes", handler.CreateNote)
    app.Put("/notes/:id", handler.UpdateNoteByID)
    app.Delete("/notes/:id", handler.DeleteNoteByID)
}
