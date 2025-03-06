package repository

import (
	"go-banking-api/internal/account/domain"

	"github.com/shopspring/decimal"
)

type AccountCommandRepositoryInterface interface {
	CreateAccount(account domain.Account)(domain.Account, error)
	UpdateAccountBalance(accountNumber string, balance decimal.Decimal)(domain.Account, error)
}

type AccountQueryRepositoryInterface interface {
	GetAccountByAccountNumber(accountNumber string) (domain.Account, error)
	GetAccountBalance(accountNumber string)(decimal.Decimal, error)
}