package database

import (
	"log"

	"gorm.io/gorm"

	ea "go-banking-api/internal/account/entity"
	ec "go-banking-api/internal/customer/entity"
	et "go-banking-api/internal/transaction/entity"
	"go-banking-api/pkg/validator"
)

func Migration(db *gorm.DB) {
	migrator := db.Migrator()

	validator.CreateEnumIfNotExists(db, "account_status_enum", "'active', 'inactive'")
	validator.CreateEnumIfNotExists(db, "transaction_type_enum", "'deposit', 'withdraw', 'none'")

	db.AutoMigrate(
		&ec.Customer{},
		&ea.Account{},
		&et.Transaction{},
	)

	tables := []string{"customers", "accounts", "transactions"}
	for _, table := range tables {
		if !migrator.HasTable(table) {
			log.Fatalf("table %s was not successfully created", table)
		}
	}
	log.Println("all tables were successfully migrated")
}
