package users

import "time"

type User struct {
	ID             int       `json:"user_id" gorm:"column:user_id"`
	Name           string    `json:"name" gorm:"column:name"`
	Occupation     string    `json:"occupation" gorm:"column:occupation"`
	Email          string    `json:"email" gorm:"column:email"`
	PasswordHash   string    `json:"password" gorm:"column:password_hash"`
	AvatarFileName string    `json:"avatar_file_name" gorm:"column:avatar_field_name"`
	Role           string    `json:"role" gorm:"column:role"`
	CreatedAt      time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"column:updated_at"`
}
