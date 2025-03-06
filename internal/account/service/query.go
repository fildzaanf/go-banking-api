package service

import (
	"errors"
	"go-banking-api/internal/account/domain"
	"go-banking-api/internal/account/repository"
)

type accountQueryService struct {
	accountQueryRepository repository.AccountQueryRepositoryInterface
}

func NewAccountQueryService(aqr repository.AccountQueryRepositoryInterface) AccountQueryServiceInterface {
	return &accountQueryService{
		accountQueryRepository: aqr,
	}
}

func (aqs *accountQueryService) GetAccountBalance(accountNumber string) (domain.Account, error) {
	if accountNumber == "" {
		return domain.Account{}, errors.New("account number not found")
	}

	account, err := aqs.accountQueryRepository.GetAccountByAccountNumber(accountNumber)
	if err != nil {
		return domain.Account{}, err
	}

	if account.AccountNumber != accountNumber {
		return domain.Account{}, errors.New("account number not found")
	}

	return account, nil
}
