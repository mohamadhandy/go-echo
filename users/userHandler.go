package users

import (
	"go-echo/helper"
	"net/http"

	"github.com/labstack/echo"
)

type userHandler struct {
	userService userService
}

func NewUserHandler(userService userService) *userHandler {
	return &userHandler{userService: userService}
}

func (u *userHandler) RegisterUser(c echo.Context) error {
	var input RegisterUserInput
	err := c.Bind(&input)
	if err != nil {
		response := helper.APIResponse("Register user failed", http.StatusUnprocessableEntity, "error", err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return err
	}
	newUser, err := u.userService.RegisterUser(input)
	if err != nil {
		res := helper.APIResponse("Register user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return err
	}
	userDTO := FormatUser(newUser, "")
	response := helper.APIResponse("Register User Success", http.StatusCreated, "success", userDTO)
	return c.JSON(http.StatusCreated, response)
}
