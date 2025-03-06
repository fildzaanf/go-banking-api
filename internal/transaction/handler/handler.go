package handler

import (
	"go-banking-api/internal/transaction/dto"
	"go-banking-api/internal/transaction/service"
	"go-banking-api/pkg/constant"
	"go-banking-api/pkg/middleware"
	"go-banking-api/pkg/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type transactionHandler struct {
	transactionQueryService   service.TransactionQueryServiceInterface
	transactionCommandService service.TransactionCommandServiceInterface
}

func NewTransactionHandler(tqs service.TransactionQueryServiceInterface, tcs service.TransactionCommandServiceInterface) *transactionHandler {
	return &transactionHandler{
		transactionQueryService:   tqs,
		transactionCommandService: tcs,
	}
}

// command
func (th *transactionHandler) CreateTransactionDeposit(c echo.Context) error {
	tokenCustomerID, errExtract := middleware.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, errExtract.Error()))
	}

	depositRequest := dto.TransactionDepositRequest{}

	if err := c.Bind(&depositRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	transactionDomain := dto.TransactionDepositRequestToDomain(depositRequest)

	CreateDeposit, err := th.transactionCommandService.CreateTransactionDeposit(transactionDomain, tokenCustomerID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	depositResponse := dto.TransactionBalanceDomainToResponse(CreateDeposit)

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "deposit completed successfully.", depositResponse))
}

func (th *transactionHandler) CreateTransactionWithdrawal(c echo.Context) error {
	tokenCustomerID, errExtract := middleware.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, errExtract.Error()))
	}

	withdrawRequest := dto.TransactionWithdrawRequest{}

	if err := c.Bind(&withdrawRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	transactionDomain := dto.TransactionWithdrawRequestToDomain(withdrawRequest)

	createdWithdrawal, err := th.transactionCommandService.CreateTransactionWithdrawal(transactionDomain, tokenCustomerID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	withdrawalResponse := dto.TransactionBalanceDomainToResponse(createdWithdrawal)

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "withdrawal completed successfully.", withdrawalResponse))
}

// query
func (th *transactionHandler) GetAllTransactions(c echo.Context) error {

	tokenCustomerID, errExtract := middleware.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, errExtract.Error()))
	}

	transactions, err := th.transactionQueryService.GetAllTransactions(tokenCustomerID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse(http.StatusUnauthorized, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, constant.SUCCESS_RETRIEVED, transactions))
}
