package dboperation

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/SystemEngineeringTeam/ryuou-manager/model"
)

func CreateUser(user model.User) error {
	db := gormConnect()
	defer db.Close()

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func SelectAllUsers() ([]model.User, error) {
	db := gormConnect()
	defer db.Close()

	var users []model.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func Login(user model.User) (model.LoginResponse, error) {
	db := gormConnect()
	defer db.Close()

	const (
		adminMail     = "admin@admin.admin"
		adminPassword = "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918"
	)

	if user.Email == adminMail && user.Password == adminPassword {
		return model.LoginResponse{
			UserID:    -1,
			TeamID:    -1,
			SessionID: "admin",
		}, nil
	}

	var member model.TeamMember
	if err := db.Model(&model.User{}).Joins("left join team_members on users.id = team_members.user_id").Where("email = ? and password = ?", user.Email, user.Password).First(&user).Scan(&member).Error; err != nil {
		return model.LoginResponse{}, err
	}

	if user.ID == 0 {
		return model.LoginResponse{}, errors.New("user not found")
	}

	var oldSession model.Session
	db.Where("user_id = ?", user.ID).First(&oldSession).Delete(oldSession)

	token := user.Email + time.Now().String()

	b := sha256.Sum256([]byte(token))
	token = hex.EncodeToString(b[:])

	session := model.Session{
		UserID:    user.ID,
		SessionID: token,
	}

	if err := db.Create(&session).Error; err != nil {
		return model.LoginResponse{}, err
	}

	return model.LoginResponse{
		UserID:    user.ID,
		TeamID:    member.TeamID,
		SessionID: token,
	}, nil
}
