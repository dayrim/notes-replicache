package handlers

import (
    "myproject/models"
    "github.com/gofiber/fiber/v2"
    "github.com/jmoiron/sqlx"
)

type NoteHandler struct {
    DB *sqlx.DB
}

func (h *NoteHandler) GetNotes(c *fiber.Ctx) error {
    notes := []models.Note{}
    err := h.DB.Select(&notes, "SELECT * FROM notes")
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.JSON(notes)
}

func (h *NoteHandler) GetNoteByID(c *fiber.Ctx) error {
    id := c.Params("id")
    note := models.Note{}
    err := h.DB.Get(&note, "SELECT * FROM notes WHERE id=$1", id)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.JSON(note)
}

func (h *NoteHandler) CreateNote(c *fiber.Ctx) error {
    note := new(models.Note)
    if err := c.BodyParser(note); err != nil {
        return c.Status(400).SendString(err.Error())
    }
    query := "INSERT INTO notes (title, content) VALUES ($1, $2) RETURNING id, created_at"
    err := h.DB.QueryRow(query, note.Title, note.Content).Scan(&note.ID, &note.CreatedAt)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.Status(201).JSON(note)
}

func (h *NoteHandler) UpdateNoteByID(c *fiber.Ctx) error {
    id := c.Params("id")
    note := new(models.Note)
    if err := c.BodyParser(note); err != nil {
        return c.Status(400).SendString(err.Error())
    }
    query := "UPDATE notes SET title=$1, content=$2 WHERE id=$3"
    _, err := h.DB.Exec(query, note.Title, note.Content, id)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.SendStatus(204)
}

func (h *NoteHandler) DeleteNoteByID(c *fiber.Ctx) error {
    id := c.Params("id")
    query := "DELETE FROM notes WHERE id=$1"
    _, err := h.DB.Exec(query, id)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.SendStatus(204)
}
