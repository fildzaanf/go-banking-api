package dto

import (
	"go-banking-api/internal/account/domain"

	"github.com/shopspring/decimal"
)

type TransactionBalanceResponse struct {
	Balance decimal.Decimal `json:"balance"`
}

// mapper
func TransactionBalanceDomainToResponse(response domain.Account) TransactionBalanceResponse {
	return TransactionBalanceResponse{
		Balance: response.Balance,
	}
}
