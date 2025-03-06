package repository

import (
	"errors"
	"go-banking-api/internal/account/domain"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type accountQueryRepository struct {
	db *gorm.DB
}

func NewAccountQueryRepository(db *gorm.DB) AccountQueryRepositoryInterface {
	return &accountQueryRepository{
		db: db,
	}
}

func (aqr *accountQueryRepository) GetAccountByAccountNumber(accountNumber string) (domain.Account, error) {
	var account domain.Account
	if err := aqr.db.Where("account_number = ?", accountNumber).First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Account{}, errors.New("account not found")
		}
		return domain.Account{}, err
	}
	return account, nil
}

func (aqr *accountQueryRepository) GetAccountBalance(accountNumber string) (decimal.Decimal, error) {
	var account domain.Account
	if err := aqr.db.Select("balance").Where("account_number = ?", accountNumber).First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return decimal.Zero, errors.New("account not found")
		}
		return decimal.Zero, err
	}
	return account.Balance, nil
}
