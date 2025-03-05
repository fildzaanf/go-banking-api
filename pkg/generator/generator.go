package generator

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateBankAccountNumber() string {
	rand.Seed(time.Now().UnixNano())

	bankCode := "008"

	branchCode := fmt.Sprintf("%04d", rand.Intn(10000))

	customerNumber := fmt.Sprintf("%06d", rand.Intn(1000000))

	checkDigit := fmt.Sprintf("%d", rand.Intn(10))

	return fmt.Sprintf("%s%s%s%s", bankCode, branchCode, customerNumber, checkDigit)
}
