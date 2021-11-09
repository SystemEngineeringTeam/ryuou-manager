package dboperation

import (
	"github.com/SystemEngineeringTeam/ryuou-manager/model"
)

func QuestionExists(id int) bool {
	db := gormConnect()
	defer db.Close()

	var question model.Question
	if err := db.Model(&question).Where("id = ?", id).First(&question).Error; err != nil {
		return false
	}

	return true
}

func TeamExists(id int) bool {
	db := gormConnect()
	defer db.Close()

	var team model.Team
	if err := db.Model(&team).Where("id = ?", id).First(&team).Error; err != nil {
		return false
	}

	return true
}

func UserExists(id int) bool {
	db := gormConnect()
	defer db.Close()

	var user model.User
	if err := db.Model(&user).Where("id = ?", id).First(&user).Error; err != nil {
		return false
	}

	return true
}
