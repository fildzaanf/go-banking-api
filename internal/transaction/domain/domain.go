package domain

import (
	"go-banking-api/internal/transaction/entity"
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID              string
	AccountID       string
	Amount          decimal.Decimal
	AccountNumber   string
	TransactionType string
	CreatedAt       time.Time
}

// mapper
func TransactionDomainToEntity(transactionDomain Transaction) entity.Transaction {
	return entity.Transaction{
		ID:              transactionDomain.ID,
		AccountID:       transactionDomain.AccountID,
		AccountNumber:   transactionDomain.AccountNumber,
		Amount:          transactionDomain.Amount,
		TransactionType: transactionDomain.TransactionType,
		CreatedAt:       transactionDomain.CreatedAt,
	}
}

func TransactionEntityToDomain(transactionEntity entity.Transaction) Transaction {
	return Transaction{
		ID:              transactionEntity.ID,
		AccountID:       transactionEntity.AccountID,
		Amount:          transactionEntity.Amount,
		AccountNumber:   transactionEntity.AccountNumber,
		TransactionType: transactionEntity.TransactionType,
		CreatedAt:       transactionEntity.CreatedAt,
	}
}

func ListTransactionDomainToEntity(transactionDomains []Transaction) []entity.Transaction {
	transactionEntities := make([]entity.Transaction, len(transactionDomains))
	for i, transaction := range transactionDomains {
		transactionEntities[i] = TransactionDomainToEntity(transaction)
	}
	return transactionEntities
}

func ListTransactionEntityToDomain(transactionEntities []entity.Transaction) []Transaction {
	transactionDomains := make([]Transaction, len(transactionEntities))
	for i, transaction := range transactionEntities {
		transactionDomains[i] = TransactionEntityToDomain(transaction)
	}
	return transactionDomains
}
