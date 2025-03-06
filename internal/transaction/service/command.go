package service

import (
	"errors"
	da "go-banking-api/internal/account/domain"
	repositoryAccount "go-banking-api/internal/account/repository"
	customerRepository "go-banking-api/internal/customer/repository"
	"go-banking-api/internal/transaction/domain"
	"go-banking-api/internal/transaction/repository"
	"go-banking-api/pkg/constant"

	"github.com/shopspring/decimal"
)

type transactionCommandService struct {
	transactionCommandRepository repository.TransactionCommandRepositoryInterface
	accountQueryRepository       repositoryAccount.AccountQueryRepositoryInterface
	accountCommandRepository     repositoryAccount.AccountCommandRepositoryInterface
	customerQueryRepository      customerRepository.CustomerQueryRepositoryInterface
}

func NewTransactionCommandService(tcr repository.TransactionCommandRepositoryInterface, aqr repositoryAccount.AccountQueryRepositoryInterface, acr repositoryAccount.AccountCommandRepositoryInterface, cqr customerRepository.CustomerQueryRepositoryInterface) TransactionCommandServiceInterface {
	return &transactionCommandService{
		transactionCommandRepository: tcr,
		accountQueryRepository:       aqr,
		accountCommandRepository:     acr,
		customerQueryRepository:      cqr,
	}
}

func (tcs *transactionCommandService) CreateTransactionDeposit(transaction domain.Transaction, customerID string) (da.Account, error) {

	if transaction.Amount.LessThanOrEqual(decimal.NewFromInt(0)) {
		return da.Account{}, errors.New("invalid deposit amount, must be greater than 0")
	}

	customer, err := tcs.customerQueryRepository.GetCustomerByID(customerID)
	if err != nil {
		return da.Account{}, err
	}

	if customer.ID != customerID {
		return da.Account{}, errors.New(constant.ERROR_ACCESS)
	}

	account, err := tcs.accountQueryRepository.GetAccountByAccountNumber(transaction.AccountNumber)
	if err != nil {
		return da.Account{}, errors.New("account number not found")
	}

	_, err = tcs.transactionCommandRepository.CreateTransaction(transaction)
	if err != nil {
		return da.Account{}, errors.New("failed to create transaction record")
	}

	account.Balance = account.Balance.Add(transaction.Amount)

	updatedBalance, err := tcs.accountCommandRepository.UpdateAccountBalance(account.AccountNumber, account.Balance)
	if err != nil {
		return da.Account{}, err
	}

	return updatedBalance, nil
}

func (tcs *transactionCommandService) CreateTransactionWithdrawal(transaction domain.Transaction, customerID string) (da.Account, error) {

	if transaction.Amount.LessThanOrEqual(decimal.NewFromInt(0)) {
		return da.Account{}, errors.New("invalid withdrawal amount, must be greater than 0")
	}

	customer, err := tcs.customerQueryRepository.GetCustomerByID(customerID)
	if err != nil {
		return da.Account{}, err
	}

	if customer.ID != customerID {
		return da.Account{}, errors.New(constant.ERROR_ACCESS)
	}

	account, err := tcs.accountQueryRepository.GetAccountByAccountNumber(transaction.AccountNumber)
	if err != nil {
		return da.Account{}, errors.New("account number not found")
	}

	if account.Balance.LessThan(transaction.Amount) {
		return da.Account{}, errors.New("insufficient balance")
	}

	_, err = tcs.transactionCommandRepository.CreateTransaction(transaction)
	if err != nil {
		return da.Account{}, errors.New("failed to create transaction record")
	}

	account.Balance = account.Balance.Sub(transaction.Amount)

	updatedBalance, err := tcs.accountCommandRepository.UpdateAccountBalance(account.AccountNumber, account.Balance)
	if err != nil {
		return da.Account{}, err
	}

	return updatedBalance, nil
}
