package domain

import (
	"go-banking-api/internal/customer/entity"
	"time"
)

type Customer struct {
	ID              string
	Name            string
	NIK             string
	PhoneNumber     string
	Password        string
	ConfirmPassword string
	CreatedAt       time.Time
}

// mapper
func CustomerDomainToEntity(customerDomain Customer) entity.Customer {
	return entity.Customer{
		ID:          customerDomain.ID,
		Name:        customerDomain.Name,
		NIK:         customerDomain.NIK,
		PhoneNumber: customerDomain.PhoneNumber,
		Password:    customerDomain.Password,
		CreatedAt:   customerDomain.CreatedAt,
	}
}

func CustomerEntityToDomain(customerEntity entity.Customer) Customer {
	return Customer{
		ID:          customerEntity.ID,
		Name:        customerEntity.Name,
		NIK:         customerEntity.NIK,
		PhoneNumber: customerEntity.PhoneNumber,
		Password:    customerEntity.Password,
		CreatedAt:   customerEntity.CreatedAt,
	}
}
