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

func Login(user model.User) (model.Session, error) {
	db := gormConnect()
	defer db.Close()

	if err := db.Where("email = ? and password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return model.Session{}, err
	}

	if user.ID == 0 {
		return model.Session{}, errors.New("user not found")
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
		return model.Session{}, err
	}

	return session, nil
}
