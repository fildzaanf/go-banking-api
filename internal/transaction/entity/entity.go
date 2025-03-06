package entity

import (
	"go-banking-api/internal/account/entity"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Transaction struct {
	ID              string          `gorm:"primaryKey"`
	AccountID       string          `gorm:"not null"`
	AccountNumber   string          `gorm:"not null"`
	Amount          decimal.Decimal `gorm:"type:decimal(15,2);not null"`
	TransactionType string          `gorm:"type:transaction_type_enum;default:'none';not null"`
	CreatedAt       time.Time       `gorm:"default:current_timestamp"`
	Account         entity.Account  `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	t.ID = UUID.String()

	if t.TransactionType == "" {
		t.TransactionType = "none"
	}

	return
}
