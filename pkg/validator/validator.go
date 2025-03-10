package validator

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
	"time"

	"gorm.io/gorm"
)

func IsDataEmpty(fields []string, data ...interface{}) error {
	if len(fields) != len(data) {
		return errors.New("column names and data length mismatch")
	}

	for i, value := range data {
		switch v := value.(type) {
		case string:
			if v == "" {
				return fmt.Errorf("%s is empty", fields[i])
			}
		case int:
			if v == 0 {
				return fmt.Errorf("%s is empty", fields[i])
			}
		case time.Time:
			if v.IsZero() {
				return fmt.Errorf("%s is empty", fields[i])
			}
		case []interface{}:
			if len(v) == 0 {
				return fmt.Errorf("%s is empty", fields[i])
			}
		case []string:
			if len(v) == 0 {
				return fmt.Errorf("%s is empty", fields[i])
			}
		case []int:
			if len(v) == 0 {
				return fmt.Errorf("%s is empty", fields[i])
			}
		default:
			if reflect.TypeOf(v).Kind() == reflect.Slice {
				slice := reflect.ValueOf(v)
				if slice.Len() == 0 {
					return fmt.Errorf("%s is empty", fields[i])
				}
			} else {
				return fmt.Errorf("unsupported data type for %s: %T", fields[i], v)
			}
		}
	}
	return nil
}

func IsEmailValid(email string) error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return errors.New("invalid email format. example: emailname@example.com")
	}
	return nil
}

func IsMinLengthValid(minLength int, fields map[string]string) error {
	for fieldName, fieldValue := range fields {
		if len(fieldValue) < minLength {
			return fmt.Errorf("minimum length for field %s is %d characters", fieldName, minLength)
		}
	}
	return nil
}

func IsMaxLengthValid(maxLength int, fields map[string]string) error {
	for fieldName, fieldValue := range fields {
		if len(fieldValue) > maxLength {
			return fmt.Errorf("maximum length for field %s is %d characters", fieldName, maxLength)
		}
	}
	return nil
}

func IsDataValid(data interface{}, validData []interface{}, caseSensitive bool) error {
	dataStr := fmt.Sprintf("%v", data)
	validDataStr := make([]string, len(validData))
	for i, v := range validData {
		validDataStr[i] = fmt.Sprintf("%v", v)
	}

	if !caseSensitive {
		dataStr = strings.ToLower(dataStr)
		for i, v := range validDataStr {
			validDataStr[i] = strings.ToLower(v)
		}
	}

	for _, validValue := range validDataStr {
		if dataStr == validValue {
			return nil
		}
	}

	return errors.New("invalid data. allowed data: " + strings.Join(validDataStr, ", "))
}

func IsDateValid(date string) error {
	if date == "" {
		return nil
	}

	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if !dateRegex.MatchString(date) {
		return errors.New("invalid date format. expected format: '2000-12-30'")
	}
	return nil
}

func IsNIKValid(nik string) error {
	nikRegex := regexp.MustCompile(`^\d{16}$`)
	if !nikRegex.MatchString(nik) {
		return errors.New("invalid NIK format. expected 16 digits")
	}
	return nil
}

func IsPhoneNumberValid(phone string) error {
	phoneRegex := regexp.MustCompile(`^\+\d{10,15}$`)
	if !phoneRegex.MatchString(phone) {
		return errors.New("invalid phone number format. expected format: +[country code] followed by 10-15 digits")
	}
	return nil
}

func CreateEnumIfNotExists(db *gorm.DB, enumName, values string) {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = '%s')", enumName)
	db.Raw(query).Scan(&exists)

	if !exists {
		createEnumSQL := fmt.Sprintf("CREATE TYPE %s AS ENUM (%s)", enumName, values)
		db.Exec(createEnumSQL)
		log.Printf("Enum %s created successfully", enumName)
	} else {
		log.Printf("Enum %s already exists, skipping creation", enumName)
	}
}
