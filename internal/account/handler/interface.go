package handler

import "github.com/labstack/echo/v4"

type UserHandlerInterface interface {
	// Query
	GetAccountBalance(c echo.Context) error
}
