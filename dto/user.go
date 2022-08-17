package dto

import "go-echo/users"

type DTOUser struct {
	ID         int    `json:"user_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	Occupation string `json:"occupation"`
	ImageURL   string `json:"image_url"`
}

func FormatUser(user users.User, token string) DTOUser {
	dtoUser := DTOUser{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Occupation: user.Occupation,
		ImageURL:   user.AvatarFileName,
		Token:      token,
	}
	return dtoUser
}
