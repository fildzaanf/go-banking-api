package entity

import (
	"go-banking-api/internal/customer/entity"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Account struct {
	ID            string          `gorm:"primaryKey"`
	CustomerID    string          `gorm:"not null"`
	AccountNumber string          `gorm:"unique;not null"`
	Balance       decimal.Decimal `gorm:"type:decimal(15,2);default:0.00;not null"`
	Status        string          `gorm:"type:account_status_enum;default:'inactive';not null"`
	CreatedAt     time.Time       `gorm:"default:current_timestamp"`
	UpdatedAt     time.Time       `gorm:"default:current_timestamp"`
	Customer      entity.Customer `gorm:"foreignKey:CustomerID;constraint:OnDelete:CASCADE"`
}

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	a.ID = UUID.String()

	if a.Status == "" {
		a.Status = "inactive"
	}

	return
}
