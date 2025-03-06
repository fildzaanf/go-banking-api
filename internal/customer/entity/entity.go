package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID          string    `gorm:"primaryKey"`
	Name        string    `gorm:"type:varchar(255);not null"`
	NIK         string    `gorm:"unique;not null"`
	PhoneNumber string    `gorm:"unique;not null"`
	Password    string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
}

func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	c.ID = UUID.String()
	return
}
