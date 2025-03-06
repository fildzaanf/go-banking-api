package dto

import (
	"go-banking-api/internal/account/domain"

	"github.com/shopspring/decimal"
)

type AccountBalanceResponse struct {
	Balance decimal.Decimal `json:"balance"`
}

// mapper
func AccountDomainToBalanceResponse(response domain.Account) AccountBalanceResponse {
	return AccountBalanceResponse{
		Balance: response.Balance,
	}
}
