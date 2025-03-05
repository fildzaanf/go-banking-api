package repository

import (
	"errors"
	"go-banking-api/internal/customer/domain"

	"gorm.io/gorm"
)

type customerQueryRepository struct {
	db *gorm.DB
}

func NewCustomerQueryRepository(db *gorm.DB) CustomerQueryRepositoryInterface {
	return &customerQueryRepository{
		db: db,
	}
}

func (cqr *customerQueryRepository) GetCustomerByNIK(nik string) (domain.Customer, error) {
	var customer domain.Customer
	result := cqr.db.Where("nik = ?", nik).First(&customer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.Customer{}, errors.New("customer not found")
		}
		return domain.Customer{}, result.Error
	}
	return customer, nil
}

func (cqr *customerQueryRepository) GetCustomerByPhoneNumber(phoneNumber string) (domain.Customer, error) {
	var customer domain.Customer
	result := cqr.db.Where("phone_number = ?", phoneNumber).First(&customer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.Customer{}, errors.New("customer not found")
		}
		return domain.Customer{}, result.Error
	}
	return customer, nil
}
