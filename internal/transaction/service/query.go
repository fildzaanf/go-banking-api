package service

import (
	"errors"
	customerRepository "go-banking-api/internal/customer/repository"
	"go-banking-api/internal/transaction/domain"
	"go-banking-api/internal/transaction/repository"
	"go-banking-api/pkg/constant"
)

type transactionQueryService struct {
	transactionQueryRepository repository.TransactionQueryRepositoryInterface
	customerQueryRepository    customerRepository.CustomerQueryRepositoryInterface
}

func NewTransactionQueryService(tqr repository.TransactionQueryRepositoryInterface, cqr customerRepository.CustomerQueryRepositoryInterface) TransactionQueryServiceInterface {
	return &transactionQueryService{
		transactionQueryRepository: tqr,
		customerQueryRepository:    cqr,
	}
}

func (tqs *transactionQueryService) GetAllTransactions(customerID string) ([]domain.Transaction, error) {
	if customerID == "" {
		return nil, errors.New(constant.ERROR_ID_INVALID)
	}

	transactions, err := tqs.transactionQueryRepository.GetAllTransactions(customerID)
	if err != nil {
		return nil, errors.New(constant.ERROR_DATA_EMPTY)
	}

	customer, err := tqs.customerQueryRepository.GetCustomerByID(customerID)
	if err != nil {
		return nil, err
	}

	if customer.ID != customerID {
		return nil, errors.New(constant.ERROR_ACCESS)
	}

	return transactions, nil
}
