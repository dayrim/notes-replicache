package handlers

import (
    "notes/models"
    "github.com/gofiber/fiber/v2"
    "github.com/jmoiron/sqlx"
)

type NoteHandler struct {
    DB *sqlx.DB
}

// GetNotes godoc
// @Summary Get all notes
// @Description Get all notes
// @Tags notes
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Note
// @Router /notes [get]
func (h *NoteHandler) GetNotes(c *fiber.Ctx) error {
    notes := []models.Note{}
    err := h.DB.Select(&notes, "SELECT * FROM notes")
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.JSON(notes)
}

// GetNoteByID godoc
// @Summary Get note by ID
// @Description Get note by ID
// @Tags notes
// @Accept  json
// @Produce  json
// @Param id path int true "Note ID"
// @Success 200 {object} models.Note
// @Router /notes/{id} [get]
func (h *NoteHandler) GetNoteByID(c *fiber.Ctx) error {
    id := c.Params("id")
    note := models.Note{}
    err := h.DB.Get(&note, "SELECT * FROM notes WHERE id=$1", id)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.JSON(note)
}

// CreateNote godoc
// @Summary Create a new note
// @Description Create a new note
// @Tags notes
// @Accept  json
// @Produce  json
// @Param note body models.Note true "New note"
// @Success 201 {object} models.Note
// @Router /notes [post]
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

// UpdateNoteByID godoc
// @Summary Update note by ID
// @Description Update note by ID
// @Tags notes
// @Accept  json
// @Produce  json
// @Param id path int true "Note ID"
// @Param note body models.Note true "Updated note"
// @Success 204
// @Router /notes/{id} [put]
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

// DeleteNoteByID godoc
// @Summary Delete note by ID
// @Description Delete note by ID
// @Tags notes
// @Accept  json
// @Produce  json
// @Param id path int true "Note ID"
// @Success 204
// @Router /notes/{id} [delete]
func (h *NoteHandler) DeleteNoteByID(c *fiber.Ctx) error {
    id := c.Params("id")
    query := "DELETE FROM notes WHERE id=$1"
    _, err := h.DB.Exec(query, id)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.SendStatus(204)
}
