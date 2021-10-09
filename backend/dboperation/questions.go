package dboperation

import (
	"github.com/SystemEngineeringTeam/ryuou-manager/model"
)

func SelectAllQuestions() ([]model.Question, error) {
	db := gormConnect()
	defer db.Close()
	var questions []model.Question
	if err := db.Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}
