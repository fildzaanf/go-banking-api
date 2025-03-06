package service

import (
	"go-banking-api/internal/account/repository"
)

type accountCommandService struct {
	accountCommandRepository repository.AccountCommandRepositoryInterface
	accountQueryRepository   repository.AccountQueryRepositoryInterface
}

func NewAccountCommandService(acr repository.AccountCommandRepositoryInterface, aqr repository.AccountQueryRepositoryInterface) AccountCommandServiceInterface {
	return &accountCommandService{
		accountCommandRepository: acr,
		accountQueryRepository:   aqr,
	}
}
