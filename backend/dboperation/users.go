package dboperation

import "github.com/SystemEngineeringTeam/ryuou-manager/model"

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
