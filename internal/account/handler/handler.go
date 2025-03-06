package handler

import (
	"go-banking-api/internal/account/dto"
	"go-banking-api/internal/account/service"
	"go-banking-api/pkg/constant"
	"go-banking-api/pkg/middleware"
	"go-banking-api/pkg/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	accountQueryService service.AccountQueryServiceInterface
}

func NewAccountHandler(aqs service.AccountQueryServiceInterface) *AccountHandler {
	return &AccountHandler{
		accountQueryService: aqs,
	}
}

// query
func (ah *AccountHandler) GetAccountBalance(c echo.Context) error {
	tokenCustomerID, errExtract := middleware.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, errExtract.Error()))
	}

	accountNumber := c.Param("account_number")
	if accountNumber == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, constant.ERROR_ID_NOT_FOUND))
	}

	account, err := ah.accountQueryService.GetAccountBalance(accountNumber)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if account.CustomerID != tokenCustomerID {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, constant.ERROR_ACCESS))
	}

	accountResponse := dto.AccountDomainToBalanceResponse(account)

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, constant.SUCCESS_RETRIEVED, accountResponse))
}
