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
