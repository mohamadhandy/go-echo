package users

import (
	"go-echo/logger"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(RegisterUserInput) (User, error)
}

type userService struct {
	userRepositoryDB UserRepositoryDB
}

func NewUserService(u UserRepositoryDB) *userService {
	return &userService{userRepositoryDB: u}
}

func (u *userService) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Email = input.Email
	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Role = input.Role

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		logger.Error("Error when encrypt password!")
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	newUser, err := u.userRepositoryDB.RegisterUser(user)
	if err != nil {
		logger.Error("Error when registering user!!")
		return user, err
	}
	return newUser, nil
}
