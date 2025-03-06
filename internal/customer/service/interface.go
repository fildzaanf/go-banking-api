package service

import "go-banking-api/internal/customer/domain"

type CustomerCommandServiceInterface interface {
	RegisterCustomer(customer domain.Customer) (string, error)
	LoginCustomer(NIK, phoneNumber, password string) (domain.Customer, string, error)
}
