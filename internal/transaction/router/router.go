package router

import (
	repositoryAccount "go-banking-api/internal/account/repository"
	repositoryCustomer "go-banking-api/internal/customer/repository"
	
	"go-banking-api/internal/transaction/handler"
	"go-banking-api/internal/transaction/repository"
	"go-banking-api/internal/transaction/service"
	"go-banking-api/pkg/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TransactionRouter(transaction *echo.Group, db *gorm.DB) {
	transactionQueryRepository := repository.NewTransactionQueryRepository(db)
	transactionCommandRepository := repository.NewTransactionCommandRepository(db)

	customerQueryRepository := repositoryCustomer.NewCustomerQueryRepository(db)

	accountQueryRepository := repositoryAccount.NewAccountQueryRepository(db)
	accountCommandRepository  := repositoryAccount.NewAccountCommandRepository(db)

	transactionCommandService := service.NewTransactionCommandService(transactionCommandRepository, accountQueryRepository, accountCommandRepository, customerQueryRepository)
	transactionQueryService := service.NewTransactionQueryService(transactionQueryRepository, customerQueryRepository)

	transactionHandler := handler.NewTransactionHandler(transactionQueryService, transactionCommandService)

	transaction.GET("", transactionHandler.GetAllTransactions, middleware.JWTMiddleware())
	transaction.POST("/deposit", transactionHandler.CreateTransactionDeposit, middleware.JWTMiddleware())
	transaction.POST("/withdrawal", transactionHandler.CreateTransactionWithdrawal, middleware.JWTMiddleware())
}
