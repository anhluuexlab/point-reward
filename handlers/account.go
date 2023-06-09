package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zett-8/go-clean-echo/logger"
	"github.com/zett-8/go-clean-echo/models"
	"github.com/zett-8/go-clean-echo/security"
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
		GetToken(c echo.Context) error
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
	sender, err := h.AccountService.GetAccountByMatID(req.SenderID)
	if err != nil {
		logger.Error("failed to account info by mattermostID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	receiver, err := h.AccountService.GetAccountByMatID(req.ReceiverID)
	if err != nil {
		logger.Error("failed to account info by mattermostID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	amount := req.Amount
	if amount < 10 {
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    utils.AmountGreatThan10.Error(),
			Data:       nil,
		})
	}

	if sender.BalanceGranted < amount {
		if sender.BalanceEarned < amount {
			return c.JSON(http.StatusBadRequest, models.Response{
				StatusCode: http.StatusBadRequest,
				Message:    utils.BalanceNotEnough.Error(),
				Data:       nil,
			})
		}
	}

	trans := &models.Transaction{
		Action:     "send",
		Amount:     req.Amount,
		SenderID:   sender.ID,
		ReceiverID: receiver.ID,
	}
	err = h.AccountService.GivePoint(sender, receiver, trans)
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
	trans, err := h.AccountService.GetTransactionByID(req.TransactionID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    utils.TransactionNotFound.Error(),
			Data:       nil,
		})
	}
	if trans.Action != "send" {
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    utils.TransactionIsRefunded.Error(),
			Data:       nil,
		})
	}
	operator, err := h.AccountService.GetAccountByMatID(req.OperatorID)
	if err != nil {
		logger.Error("failed to account info by mattermostID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	sender, err := h.AccountService.GetAccountByID(trans.SenderID)
	if err != nil {
		logger.Error("failed to account info by mattermostID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	receiver := operator
	if trans.ReceiverID != operator.ID {
		receiver, err = h.AccountService.GetAccountByID(trans.ReceiverID)
		if err != nil {
			logger.Error("failed to account info by mattermostID", zap.Error(err))
			return c.JSON(http.StatusBadRequest, models.Response{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
				Data:       nil,
			})
		}
	}
	amount := trans.Amount
	if receiver.BalanceGranted < amount {
		if receiver.BalanceEarned < amount {
			return c.JSON(http.StatusBadRequest, models.Response{
				StatusCode: http.StatusBadRequest,
				Message:    utils.BalanceNotEnough.Error(),
				Data:       nil,
			})
		}
	}
	trans.Action = "reject"
	trans.OperatorID = operator.ID

	err = h.AccountService.RejectPoint(receiver, sender, trans)
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
	requester, err := h.AccountService.GetAccountByMatID(req.RequesterID)
	if err != nil {
		logger.Error("failed to account info by mattermostID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	amount := req.Amount
	if requester.BalanceEarned < amount {
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    utils.BalanceNotEnough.Error(),
			Data:       nil,
		})
	}
	exRequest := &models.ExchangeRequests{
		Status:      "request",
		Amount:      amount,
		RequesterID: requester.ID,
	}
	err = h.AccountService.ExchangePointRequest(requester, exRequest)
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

func (h *accountHandler) GetToken(c echo.Context) error {
	token, err := security.GenToken()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       token,
	})
}
