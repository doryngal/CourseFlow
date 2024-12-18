package models

import "time"

type Review struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	CourseID  int       `json:"course_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
