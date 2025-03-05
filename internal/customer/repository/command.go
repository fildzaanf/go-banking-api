package repository

import (
	"go-banking-api/internal/customer/domain"
	"go-banking-api/internal/customer/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type customerCommandRepository struct {
	db *gorm.DB
}

func NewCustomerCommandRepository(db *gorm.DB) CustomerCommandRepositoryInterface {
	return &customerCommandRepository{
		db: db,
	}
}

func (ccr *customerCommandRepository) RegisterCustomer(customer domain.Customer) (domain.Customer, error) {

	tx := ccr.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return domain.Customer{}, tx.Error
	}

	customerEntity := domain.CustomerDomainToEntity(customer)

	if err := tx.Create(&customerEntity).Error; err != nil {
		tx.Rollback()
		return domain.Customer{}, err
	}

	customerDomain := domain.CustomerEntityToDomain(customerEntity)

	if err := tx.Commit().Error; err != nil {
		return domain.Customer{}, err
	}

	return customerDomain, nil
}
func (ccr *customerCommandRepository) LoginCustomer(NIK, phoneNumber, password string) (domain.Customer, error) {
	var customerEntity entity.Customer

	tx := ccr.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return domain.Customer{}, tx.Error
	}

	result := tx.Where("nik = ? AND phone_number = ?", NIK, phoneNumber).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&customerEntity)
	if result.Error != nil {
		tx.Rollback()
		return domain.Customer{}, result.Error
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return domain.Customer{}, err
	}

	customerDomain := domain.CustomerEntityToDomain(customerEntity)

	return customerDomain, nil
}
