package route

import (
	"cleango-template/internal/handler"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo, userHandler *handler.UserHandler) {
	e.GET("/users", userHandler.GetAllUsers)
}