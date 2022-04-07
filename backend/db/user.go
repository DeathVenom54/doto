package db

import (
	"github.com/DeathVenom54/doto-backend/snowflake"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	Email    string `gorm:"uniqueIndex"`
	Password string
}

func CreateUser(user *User) error {
	user.ID = snowflake.GenerateUint()

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return err
	}
	user.Password = string(hash)

	result := DB.Create(&user)
	return result.Error
}
