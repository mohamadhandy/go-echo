package users

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RegisterUserInput struct {
	Email      string `json:"email" binding:"required,email"`
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Role       string `json:"role" binding:"required"`
}
