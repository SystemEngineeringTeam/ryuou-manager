package dboperation

import (
	"github.com/SystemEngineeringTeam/ryuou-manager/model"
)

func SelectAllQuestionsByTeamID(teamID int) ([]model.QuestionResponse, error) {
	db := gormConnect()
	defer db.Close()

	// Hi,Mr.Copilot.
	// Could you convert this raw SQL to a gorm method?
	// I'm not sure how to do that.

	var questions []model.QuestionResponse

	if err := db.Raw(`select questions.id,questions.title,if(id=question_id,questions.description,'Not Opened') as description,if(id=question_id,1,0) as is_opened,is_passed from questions,team_opened_questions where team_id=?;`, teamID).Scan(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func SelectQuestionDetailByTeamID(teamID int, questionID int) (model.QuestionDetailResponse, error) {
	db := gormConnect()
	defer db.Close()

	var question model.QuestionDetailResponse

	if err := db.Model(&model.Question{}).Joins("join team_opened_questions on team_opened_questions.question_id=questions.id").Where("team_id=? and question_id=?", teamID, questionID).Scan(&question).Error; err != nil {
		return question, err
	}

	return question, nil
}

func OpenQuestion(teamID, questionID int) error {
	db := gormConnect()
	defer db.Close()

	if err := db.Create(&model.TeamOpenedQuestion{TeamID: teamID, QuestionID: questionID}).Error; err != nil {
		return err
	}

	return nil
}

func SubmitQuestion(teamID, questionID int, answer string) error {
	db := gormConnect()
	defer db.Close()

	if err := db.Create(&model.TeamSubmittedQuestion{TeamID: teamID, QuestionID: questionID, URL: answer}).Error; err != nil {
		return err
	}

	return nil
}
