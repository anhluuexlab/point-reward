package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zett-8/go-clean-echo/logger"
	"github.com/zett-8/go-clean-echo/models"
	"github.com/zett-8/go-clean-echo/services"
	"github.com/zett-8/go-clean-echo/utils"
	"go.uber.org/zap"
)

type (
	AccountHandler interface {
		MyBalance(c echo.Context) error
		MyTransactions(c echo.Context) error
		GivePoint(c echo.Context) error
		RejectPoint(c echo.Context) error
		ExchangeRequest(c echo.Context) error
	}

	accountHandler struct {
		services.AccountService
	}
)

func (h *accountHandler) MyBalance(c echo.Context) error {
	mattermostID := c.Param("user_id")

	account, err := h.AccountService.GetAccountByMatID(mattermostID)
	if err != nil {
		logger.Error("failed to account info by mattermostID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       account,
	})
}

func (h *accountHandler) MyTransactions(c echo.Context) error {
	mattermostID := c.Param("user_id")
	account, err := h.AccountService.GetAccountByMatID(mattermostID)
	if err != nil {
		logger.Error("failed to account info by mattermostID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	paging := utils.Paging{}
	if err := c.Bind(&paging); err != nil {
		logger.Error("lỗi nhập liệu", zap.Error(err))
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	r, err := h.AccountService.GetTransactionByMatID(account.ID, &paging)

	if err != nil {
		logger.Error("failed to get book", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       r,
	})
}

func (h *accountHandler) GivePoint(c echo.Context) error {
	req := models.GivePointForm{}
	if err := c.Bind(&req); err != nil {
		logger.Error("lỗi nhập liệu", zap.Error(err))
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	if err := c.Validate(req); err != nil {
		logger.Error("lỗi nhập liệu", zap.Error(err))
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "lỗi nhập liệu",
			Data:       err.Error(),
		})
	}
	trans := &models.Transaction{
		Action:     "send",
		Amount:     req.Amount,
		SenderID:   req.SenderID,
		ReceiverID: req.ReceiverID,
	}
	err := h.AccountService.GivePoint(trans)
	if err != nil {
		logger.Error("failed to Give Point", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       nil,
	})
}

func (h *accountHandler) RejectPoint(c echo.Context) error {
	req := models.RejectPointForm{}
	if err := c.Bind(&req); err != nil {
		logger.Error("lỗi nhập liệu", zap.Error(err))
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	if err := c.Validate(req); err != nil {
		logger.Error("lỗi nhập liệu", zap.Error(err))
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "lỗi nhập liệu",
			Data:       err.Error(),
		})
	}
	trans := &models.Transaction{
		Action:     "reject",
		ID:         req.TransactionID,
		OperatorID: req.OperatorID,
	}
	err := h.AccountService.RejectPoint(trans)
	if err != nil {
		logger.Error("failed to Reject Point", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       nil,
	})
}
func (h *accountHandler) ExchangeRequest(c echo.Context) error {
	req := models.ExchangeRequestForm{}
	if err := c.Bind(&req); err != nil {
		logger.Error("lỗi nhập liệu", zap.Error(err))
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	if err := c.Validate(req); err != nil {
		logger.Error("lỗi nhập liệu", zap.Error(err))
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "lỗi nhập liệu",
			Data:       err.Error(),
		})
	}
	exRequest := &models.ExchangeRequests{
		Status:      "request",
		Amount:      req.Amount,
		RequesterID: req.RequesterID,
	}
	err := h.AccountService.ExchangePointRequest(exRequest)
	if err != nil {
		logger.Error("failed to Exchange Request", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       nil,
	})
}
