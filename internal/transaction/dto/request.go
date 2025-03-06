package dto

import (
	"go-banking-api/internal/transaction/domain"

	"github.com/shopspring/decimal"
)

type TransactionDepositRequest struct {
	AccountNumber string          `json:"account_number" form:"account_number"`
	Amount        decimal.Decimal `json:"amount" form:"amount"`
}

type TransactionWithdrawRequest struct {
	AccountNumber string          `json:"account_number" form:"account_number"`
	Amount        decimal.Decimal `json:"amount" form:"amount"`
}

// mapper
func TransactionDepositRequestToDomain(request TransactionDepositRequest) domain.Transaction {
	return domain.Transaction{
		AccountNumber:   request.AccountNumber,
		Amount:          request.Amount,
		TransactionType: "deposit",
	}
}

func TransactionWithdrawRequestToDomain(request TransactionWithdrawRequest) domain.Transaction {
	return domain.Transaction{
		AccountNumber:   request.AccountNumber,
		Amount:          request.Amount,
		TransactionType: "withdraw",
	}
}
