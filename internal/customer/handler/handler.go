package handler

import (
	"go-banking-api/internal/customer/dto"
	"go-banking-api/internal/customer/service"
	"go-banking-api/pkg/constant"
	"go-banking-api/pkg/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type customerHandler struct {
	customerCommandService service.CustomerCommandServiceInterface
}

func NewCustomerHandler(ccs service.CustomerCommandServiceInterface) *customerHandler {
	return &customerHandler{
		customerCommandService: ccs,
	}
}

// Command
func (ch *customerHandler) RegisterCustomer(c echo.Context) error {
	customerRequest := dto.CustomerRegisterRequest{}

	errBind := c.Bind(&customerRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, errBind.Error()))
	}

	customerDomain := dto.CustomerRegisterRequestToDomain(customerRequest)

	registeredCustomer, errRegister := ch.customerCommandService.RegisterCustomer(customerDomain)
	if errRegister != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, errRegister.Error()))
	}

	customerResponse := dto.CustomerRegisterDomainToResponse(registeredCustomer)

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, constant.SUCCESS_REGISTER, customerResponse))
}

func (ch *customerHandler) LoginCustomer(c echo.Context) error {
	customerRequest := dto.CustomerLoginRequest{}

	errBind := c.Bind(&customerRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, errBind.Error()))
	}

	loginCustomer, token, errLogin := ch.customerCommandService.LoginCustomer(customerRequest.NIK, customerRequest.PhoneNumber, customerRequest.Password)
	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, errLogin.Error()))
	}

	customerResponse := dto.CustomerDomainToLoginResponse(loginCustomer, token)

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, constant.SUCCESS_LOGIN, customerResponse))
}
