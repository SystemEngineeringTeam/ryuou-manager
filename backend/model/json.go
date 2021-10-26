package model

type QuestionResponse struct {
	ID          int    `json:"id" gorm:"id"`
	Title       string `json:"title" gorm:"title"`
	Description string `json:"description" gorm:"description"`
	IsOpened    bool   `json:"isOpened" gorm:"is_opened"`
	IsPassed    bool   `json:"isPassed" gorm:"is_passed"`
}

type QuestionDetailResponse struct {
	ID          int    `json:"id" gorm:"id"`
	Title       string `json:"title" gorm:"title"`
	Description string `json:"description" gorm:"description"`
	Score       int    `json:"score" gorm:"score"`
	IsPassed    bool   `json:"isPassed" gorm:"is_passed"`
}

type SubmitResponse struct {
	ID         int    `json:"id" gorm:"id"`
	Title      string `json:"title" gorm:"title"`
	TeamID     int    `json:"teamID" gorm:"team_id"`
	QuestionID int    `json:"questionID" gorm:"question_id"`
	URL        string `json:"url" gorm:"url"`
}

type SubmitRequest struct {
	Answer string `json:"answer"`
}
