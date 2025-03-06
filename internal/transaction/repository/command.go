package repository

import (
	"errors"
	"go-banking-api/internal/account/entity"
	"go-banking-api/internal/transaction/domain"

	"gorm.io/gorm"
)

type transactionCommandRepository struct {
	db *gorm.DB
}

func NewTransactionCommandRepository(db *gorm.DB) TransactionCommandRepositoryInterface {
	return &transactionCommandRepository{
		db: db,
	}
}

func (r *transactionCommandRepository) CreateTransaction(transaction domain.Transaction) (domain.Transaction, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return domain.Transaction{}, tx.Error
	}

	var account entity.Account
	err := tx.Raw("SELECT * FROM accounts WHERE account_number = ? FOR UPDATE", transaction.AccountNumber).
		Scan(&account).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Transaction{}, errors.New("account not found")
		}
		return domain.Transaction{}, err
	}

	transaction.AccountID = account.ID

	transactionEntity := domain.TransactionDomainToEntity(transaction)

	if err := tx.Create(&transactionEntity).Error; err != nil {
		tx.Rollback()
		return domain.Transaction{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return domain.Transaction{}, err
	}

	transactionDomain := domain.TransactionEntityToDomain(transactionEntity)

	return transactionDomain, nil
}
