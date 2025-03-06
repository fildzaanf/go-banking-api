package router

import (
	"go-banking-api/pkg/middleware"
	"go-banking-api/internal/account/repository"
	"go-banking-api/internal/account/service"
	"go-banking-api/internal/account/handler"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AccountRouter(account *echo.Group, db *gorm.DB) {
	accountQueryRepository := repository.NewAccountQueryRepository(db)

	accountQueryService := service.NewAccountQueryService(accountQueryRepository)

	accountHandler := handler.NewAccountHandler(accountQueryService)

	account.GET("/balance/:account_number", accountHandler.GetAccountBalance, middleware.JWTMiddleware())
}
