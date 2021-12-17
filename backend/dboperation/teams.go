package dboperation

import (
	"log"

	"github.com/SystemEngineeringTeam/ryuou-manager/model"
)

func SelectAllTeams() []model.Team {
	db := gormConnect()
	defer db.Close()

	var teams []model.Team
	if err := db.Find(&teams).Error; err != nil {
		log.Fatal(err)
	}

	return teams
}

func JoinTeam(teamID, userID int) error {
	db := gormConnect()
	defer db.Close()

	member := model.TeamMember{
		UserID: userID,
		TeamID: teamID,
	}

	if err := db.Create(&member).Error; err != nil {
		return err
	}
	return nil
}

func LeaveTeam(teamID, userID int) error {
	db := gormConnect()
	defer db.Close()

	if err := db.Debug().Model(&model.TeamMember{}).Where("team_id = ? and user_id = ?", teamID, userID).Delete(&model.TeamMember{}).Error; err != nil {
		return err
	}
	return nil
}

func AddScore(questionID, teamID int) error {
	db := gormConnect()
	defer db.Close()

	var question model.Question
	var team model.Team
	db.Model(&model.Question{}).Where("id=?", questionID).Scan(&question).Model(&model.Team{}).Where("id=?", teamID).Scan(&team).Update("score", team.Score+question.Score)

	return nil
}

func InsertNewTeam(name string) error {
	db := gormConnect()
	defer db.Close()

	team := model.Team{
		Name: name,
	}

	if err := db.Create(&team).Error; err != nil {
		return err
	}
	return nil
}

func RemoveTeam(teamID int) error {
	db := gormConnect()
	defer db.Close()

	team := model.Team{
		ID: teamID,
	}

	if err := db.Model(&model.Team{}).Where("id=?", teamID).Delete(&team).Error; err != nil {
		return err
	}
	return nil
}
