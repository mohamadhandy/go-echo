package users

import (
	"go-echo/logger"

	"gorm.io/gorm"
)

type UserRepositoryDB interface {
	RegisterUser(User) (User, error)
}

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) *userRepositoryDB {
	return &userRepositoryDB{db}
}

func (u *userRepositoryDB) RegisterUser(user User) (User, error) {
	var err error
	if err = u.db.Create(&user).Error; err != nil {
		logger.Error("Error when registering user")
		return user, err
	}
	return user, nil
}
