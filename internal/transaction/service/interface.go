package service

import (
	"go-banking-api/internal/transaction/domain"
	da "go-banking-api/internal/account/domain"
)

type TransactionCommandServiceInterface interface {
	CreateTransactionDeposit(transaction domain.Transaction, customerID string) (da.Account, error)
	CreateTransactionWithdrawal(transaction domain.Transaction, customerID string) (da.Account, error)
}

type TransactionQueryServiceInterface interface {
	GetAllTransactions(customerID string) ([]domain.Transaction, error)
}
