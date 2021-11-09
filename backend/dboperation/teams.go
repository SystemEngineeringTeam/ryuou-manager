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
