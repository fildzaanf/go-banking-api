package dto

import "go-banking-api/internal/customer/domain"

type CustomerRegisterRequest struct {
	Name            string `json:"name" form:"name"`
	NIK             string  `json:"nik" form:"nik"`
	PhoneNumber     string  `json:"phone_number" form:"phone_number"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type CustomerLoginRequest struct {
	NIK         string  `json:"nik" form:"nik"`
	PhoneNumber string  `json:"phone_number" form:"phone_number"`
	Password    string `json:"password" form:"password"`
}

// mapper
func CustomerRegisterRequestToDomain(request CustomerRegisterRequest) domain.Customer {
	return domain.Customer{
		Name:            request.Name,
		NIK:             request.NIK,
		PhoneNumber:     request.PhoneNumber,
		Password:        request.Password,
		ConfirmPassword: request.ConfirmPassword,
	}
}

func CustomerLoginRequestToDomain(request CustomerLoginRequest) domain.Customer {
	return domain.Customer{
		NIK:         request.NIK,
		PhoneNumber: request.PhoneNumber,
		Password:    request.Password,
	}
}
