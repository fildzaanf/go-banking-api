package domain

import (
	"go-banking-api/internal/account/entity"
	"go-banking-api/internal/customer/domain"
	"time"

	"github.com/shopspring/decimal"
)

type Account struct {
	ID            string
	CustomerID    string
	AccountNumber string
	Balance       decimal.Decimal
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Customer      domain.Customer
}

// mapper
func AccountDomainToEntity(accountDomain Account) entity.Account {
	return entity.Account{
		ID:            accountDomain.ID,
		CustomerID:    accountDomain.CustomerID,
		AccountNumber: accountDomain.AccountNumber,
		Balance:       accountDomain.Balance,
		Status:        accountDomain.Status,
		CreatedAt:     accountDomain.CreatedAt,
		UpdatedAt:     accountDomain.UpdatedAt,
	}
}

func AccountEntityToDomain(accountEntity entity.Account) Account {
	return Account{
		ID:            accountEntity.ID,
		CustomerID:    accountEntity.CustomerID,
		AccountNumber: accountEntity.AccountNumber,
		Balance:       accountEntity.Balance,
		Status:        accountEntity.Status,
		CreatedAt:     accountEntity.CreatedAt,
		UpdatedAt:     accountEntity.UpdatedAt,
	}
}
