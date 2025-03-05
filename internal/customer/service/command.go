package service

import (
	"errors"
	"go-banking-api/internal/customer/domain"
	da "go-banking-api/internal/account/domain"
	"go-banking-api/internal/customer/repository"
	repositoryAccount 	"go-banking-api/internal/account/repository"
	"go-banking-api/pkg/constant"
	"go-banking-api/pkg/crypto"
	"go-banking-api/pkg/generator"
	"go-banking-api/pkg/middleware"
	"go-banking-api/pkg/validator"
)

type customerCommandService struct {
	customerCommandRepository repository.CustomerCommandRepositoryInterface
	customerQueryRepository   repository.CustomerQueryRepositoryInterface
	accountCommandRepository repositoryAccount.AccountCommandRepositoryInterface
}

func NewCustomerCommandService(ccr repository.CustomerCommandRepositoryInterface, cqr repository.CustomerQueryRepositoryInterface, acr repositoryAccount.AccountCommandRepositoryInterface) CustomerCommandServiceInterface {
	return &customerCommandService{
		customerCommandRepository: ccr,
		customerQueryRepository:   cqr,
		accountCommandRepository: acr,
	}
}

func (ccs *customerCommandService) RegisterCustomer(customer domain.Customer) (string, error) {
	errEmpty := validator.IsDataEmpty([]string{"name", "password", "confirm_password", "nik", "phone_number"}, customer.Name, customer.Password, customer.ConfirmPassword, customer.NIK, customer.PhoneNumber)
	if errEmpty != nil {
		return "", errEmpty
	}

	errNIKValid := validator.IsNIKValid(customer.NIK)
	if errNIKValid != nil {
		return "", errNIKValid
	}

	errPhoneValid := validator.IsPhoneNumberValid(customer.PhoneNumber)
	if errPhoneValid != nil {
		return "", errPhoneValid
	}

	errLength := validator.IsMinLengthValid(10, map[string]string{"password": customer.Password})
	if errLength != nil {
		return "", errLength
	}

	if customer.Password != customer.ConfirmPassword {
		return "", errors.New(constant.ERROR_PASSWORD_CONFIRM)
	}

	hashedPassword, errHash := crypto.HashPassword(customer.Password)
	if errHash != nil {
		return "", errors.New(constant.ERROR_PASSWORD_HASH)
	}

	customer.Password = hashedPassword

	customerEntity, errRegister := ccs.customerCommandRepository.RegisterCustomer(customer)
	if errRegister != nil {
		return "", errRegister
	}

	accountNumber := generator.GenerateBankAccountNumber()


	account := da.Account{
		CustomerID:    customerEntity.ID,
		AccountNumber: accountNumber,
	}

	errCreateAccount := ccs.accountCommandRepository.CreateAccount(account)
	if errCreateAccount != nil {
		return "", errCreateAccount
	}

	return accountNumber, nil
}

func (ccs *customerCommandService) LoginCustomer(NIK, phoneNumber string, password string) (domain.Customer, string, error) {
	errEmpty := validator.IsDataEmpty([]string{"nik", "phone_number", "password"}, NIK, phoneNumber, password)
	if errEmpty != nil {
		return domain.Customer{}, "", errEmpty
	}

	errNIKValid := validator.IsNIKValid(NIK)
	if errNIKValid != nil {
		return domain.Customer{}, "", errNIKValid
	}

	customerDomain, errGetCustomerNIK := ccs.customerQueryRepository.GetCustomerByNIK(NIK)
	if errGetCustomerNIK == nil {
		return domain.Customer{},"", errors.New("NIK already registered")
	}

	errPhoneValid := validator.IsPhoneNumberValid(phoneNumber)
	if errPhoneValid != nil {
		return domain.Customer{}, "", errPhoneValid
	}

	customerDomain, errGetCustomerPhone := ccs.customerQueryRepository.GetCustomerByPhoneNumber(phoneNumber)
    if errGetCustomerPhone == nil {
        return domain.Customer{}, "", errors.New("phone number already registered")
    }

	comparePassword := crypto.ComparePassword(customerDomain.Password, password)
	if comparePassword != nil {
		return domain.Customer{}, "", errors.New(constant.ERROR_LOGIN)
	}

	token, errCreate := middleware.GenerateToken(customerDomain.ID)
	if errCreate != nil {
		return domain.Customer{}, "", errors.New(constant.ERROR_TOKEN_GENERATE)
	}

	return customerDomain, token, nil
}
