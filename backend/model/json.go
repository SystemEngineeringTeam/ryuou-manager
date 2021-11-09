package model

type QuestionResponse struct {
	ID          int    `json:"id" gorm:"id"`
	Title       string `json:"title" gorm:"title"`
	Description string `json:"description" gorm:"description"`
	IsOpened    bool   `json:"is_opened" gorm:"is_opened"`
	IsPassed    bool   `json:"is_passed" gorm:"is_passed"`
}

type QuestionDetailResponse struct {
	ID          int    `json:"id" gorm:"id"`
	Title       string `json:"title" gorm:"title"`
	Description string `json:"description" gorm:"description"`
	Score       int    `json:"score" gorm:"score"`
	IsPassed    bool   `json:"is_passed" gorm:"is_passed"`
}

type SubmitResponse struct {
	ID         int    `json:"id" gorm:"id"`
	Title      string `json:"title" gorm:"title"`
	TeamID     int    `json:"team_id" gorm:"team_id"`
	QuestionID int    `json:"question_id" gorm:"question_id"`
	URL        string `json:"url" gorm:"url"`
}

type SubmitRequest struct {
	Answer string `json:"answer"`
}
