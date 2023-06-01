package handlers

import (
	"net/http"
	"strconv"

	"github.com/buniekbua/gousers/models"
	"github.com/buniekbua/gousers/repositories"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userRepo *repositories.UserRepository
}

func NewUserHandler(userRepo *repositories.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

func (uh *UserHandler) CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	err := uh.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) GetUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := uh.userRepo.GetUserByID(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)

}

func (uh *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := uh.userRepo.GetAllUsers()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (uh *UserHandler) UpdateUser(c echo.Context) error {

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	errRepo := uh.userRepo.UpdateUser(userID, user)
	if errRepo != nil {
		return errRepo
	}

	return c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) DeleteUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	errRepo := uh.userRepo.DeleteUser(userID)
	if errRepo != nil {
		return errRepo
	}

	return c.JSON(http.StatusOK, "User has been deleted.")
}
