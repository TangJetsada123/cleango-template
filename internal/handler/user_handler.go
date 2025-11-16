package handler

import (
	"context"
	"net/http"
	"time"

	"cleango-template/internal/domain/model"
	"cleango-template/internal/repository"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := h.repo.FindAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: echo.Map{"data": err.Error()}})

	}

	return c.JSON(http.StatusOK, model.UserResponse{Status: http.StatusOK, Message: "success", Data: echo.Map{"data": results}})

}
