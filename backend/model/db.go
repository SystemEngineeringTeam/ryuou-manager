package model

import "time"

type Question struct {
	ID          int       `gorm:"id"`
	Title       string    `gorm:"title"`
	Description string    `gorm:"description"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
}

type TeamOpenedQuestion struct {
	TeamID     int       `gorm:"team_id"`
	QuestionID int       `gorm:"question_id"`
	IsPassed   bool      `gorm:"is_passed"`
	CreatedAt  time.Time `gorm:"created_at"`
	PassedAt   time.Time `gorm:"passed_at"`
}

type Team struct {
	ID        int       `gorm:"id"`
	Name      string    `gorm:"name"`
	Score     int       `gorm:"score"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

type TeamSubmittedQuestion struct {
	ID          int    `gorm:"id"`
	TeamID      int    `gorm:"team_id"`
	QuestionID  int    `gorm:"question_id"`
	URL         string `gorm:"url"`
	SubmittedAt int    `gorm:"submitted_at"`
}
