package models

import "time"

type Module struct {
	ID        int       `json:"id"`
	CourseID  int       `json:"course_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Order     int       `json:"order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
