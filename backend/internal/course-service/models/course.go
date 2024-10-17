package models

import "time"

type Course struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Author    string    `json:"author"`
	Duration  int       `json:"duration"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
