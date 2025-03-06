package repository

import "go-banking-api/internal/transaction/domain"

type TransactionCommandRepositoryInterface interface {
	CreateTransaction(transaction domain.Transaction)(domain.Transaction, error)
}

type TransactionQueryRepositoryInterface interface {
	GetAllTransactions(customerID string) ([]domain.Transaction, error)
}