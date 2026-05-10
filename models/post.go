package models

import "time"

type Post struct {
	ID          int       `json:"id"           db:"id"`
	Title       string    `json:"title"        db:"title"`
	Content     string    `json:"content"      db:"content"`
	Category    string    `json:"category"     db:"category"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
	UpdatedDate time.Time `json:"updated_date" db:"updated_date"`
	Status      string    `json:"status"       db:"status"`
}
