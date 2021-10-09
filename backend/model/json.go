package model

type QuestionResponse struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsOpened    bool   `json:"is_opened"`
	IsCorrect   bool   `json:"is_correct"`
}
