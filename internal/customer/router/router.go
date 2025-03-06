package router

import (
	"go-banking-api/internal/customer/handler"
	"go-banking-api/internal/customer/repository"
	"go-banking-api/internal/customer/service"
	repositoryAccount 	"go-banking-api/internal/account/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CustomerRouter(customer *echo.Group, db *gorm.DB) {
	customerQueryRepository := repository.NewCustomerQueryRepository(db)
	customerCommandRepository := repository.NewCustomerCommandRepository(db)
	accountCommandRepository := repositoryAccount.NewAccountCommandRepository(db)

	customerCommandService := service.NewCustomerCommandService(customerCommandRepository, customerQueryRepository, accountCommandRepository)

	customerHandler := handler.NewCustomerHandler(customerCommandService)

	customer.POST("/register", customerHandler.RegisterCustomer)
	customer.POST("/login", customerHandler.LoginCustomer)
}
