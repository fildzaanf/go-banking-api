package repository

import (
	"errors"
	"fmt"
	"go-banking-api/internal/account/domain"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type accountCommandRepository struct {
	db *gorm.DB
}

func NewAccountCommandRepository(db *gorm.DB) AccountCommandRepositoryInterface {
	return &accountCommandRepository{
		db: db,
	}
}

func (acr *accountCommandRepository) CreateAccount(account domain.Account) (domain.Account, error) {
	tx := acr.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return domain.Account{}, tx.Error
	}

	if err := tx.Create(&account).Error; err != nil {
		tx.Rollback()
		return domain.Account{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return domain.Account{}, err
	}

	return account, nil
}

func (acr *accountCommandRepository) UpdateAccountBalance(accountNumber string, balance decimal.Decimal) (domain.Account, error) {
	tx := acr.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return domain.Account{}, fmt.Errorf("failed to start transaction: %w", tx.Error)
	}

	var account domain.Account

	if err := tx.Raw("SELECT * FROM accounts WHERE account_number = ? FOR UPDATE", accountNumber).Scan(&account).Error; err != nil {
		tx.Rollback()
		return domain.Account{}, fmt.Errorf("failed to fetch account: %w", err)
	}

	if account.ID == "" {
		tx.Rollback()
		return domain.Account{}, errors.New("account not found")
	}

	newBalance := balance

	if newBalance.LessThan(decimal.NewFromInt(0)) {
		tx.Rollback()
		return domain.Account{}, errors.New("insufficient funds")
	}

	account.Balance = newBalance

	if err := tx.Save(&account).Error; err != nil {
		tx.Rollback()
		return domain.Account{}, fmt.Errorf("failed to update balance: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return domain.Account{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return account, nil
}
