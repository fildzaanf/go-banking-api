package handler

import "github.com/labstack/echo/v4"

type TransactionHandlerInterface interface {
	// query
	GetAllTransactions(c echo.Context)error

	// command
	CreateTransactionDeposit(c echo.Context) error
	CreateTransactionWithdrawal(c echo.Context) error
}
