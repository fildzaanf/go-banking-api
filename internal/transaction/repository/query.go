package repository

import (
	"errors"
	"go-banking-api/internal/transaction/domain"
	"go-banking-api/internal/transaction/entity"

	"gorm.io/gorm"
)

type transactionQueryRepository struct {
	db *gorm.DB
}

func NewTransactionQueryRepository(db *gorm.DB) TransactionQueryRepositoryInterface {
	return &transactionQueryRepository{
		db: db,
	}
}

func (tqr *transactionQueryRepository) GetAllTransactions(customerID string) ([]domain.Transaction, error) {
	var transactions []entity.Transaction

	result := tqr.db.
		Joins("JOIN accounts ON accounts.account_number = transactions.account_number").
		Where("accounts.customer_id = ?", customerID).
		Select("transactions.*").
		Find(&transactions)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("no transactions found for this customer")
		}
		return nil, result.Error
	}

	return domain.ListTransactionEntityToDomain(transactions), nil
}
