package repository

import (
	"go-banking-api/internal/customer/domain"
)

type CustomerCommandRepositoryInterface interface {
	RegisterCustomer(customer domain.Customer) (domain.Customer, error)
	LoginCustomer(NIK, phoneNumber, password string) (domain.Customer, error)
}

type CustomerQueryRepositoryInterface interface {
	GetCustomerByNIK(nik string) (domain.Customer, error)
	GetCustomerByPhoneNumber(phoneNumber string) (domain.Customer, error)
}
