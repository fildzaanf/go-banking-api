package service

import "go-banking-api/internal/account/domain"

type AccountCommandServiceInterface interface {
}

type AccountQueryServiceInterface interface{
	GetAccountBalance(accountNumber string)(domain.Account, error)
}
