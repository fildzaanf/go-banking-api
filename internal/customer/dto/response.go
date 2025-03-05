package dto

import "go-banking-api/internal/customer/domain"

type CustomerRegisterResponse struct {
	AccountNumber string `json:"account_number"`
}

type CustomerLoginResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

// mapper
func CustomerRegisterDomainToResponse(accountNumber string) CustomerRegisterResponse {
	return CustomerRegisterResponse{
		AccountNumber: accountNumber,
	}
}

func CustomerDomainToLoginResponse(response domain.Customer, token string) CustomerLoginResponse {
	return CustomerLoginResponse{
		ID:    response.ID,
		Token: token,
	}
}
