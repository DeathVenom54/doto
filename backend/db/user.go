package db

import (
	"github.com/DeathVenom54/doto-backend/snowflake"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	Email    string `gorm:"uniqueIndex"`
	Password string
}

func CreateUser(user *User) (*User, error) {
	user.ID = snowflake.GenerateUint()

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return nil, err
	}
	user.Password = string(hash)

	result := DB.Create(&user)
	return user, result.Error
}

func GetUserById(id uint) (*User, error) {
	user := &User{ID: id}
	res := DB.First(user)

	return user, res.Error
}
