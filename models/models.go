package models

import (
	"gin/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUser(username string) (User, error) {
	var user User
	database.Db.Where("username = ?", username).First(&user)
	return user, nil
}
