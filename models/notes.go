package models

import "time"

type Note struct {
    ID        int       `db:"id" json:"id"`
    Title     string    `db:"title" json:"title"`
    Content   string    `db:"content" json:"content"`
    CreatedAt time.Time `db:"created_at" json:"created_at"`
}
