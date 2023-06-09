package services

import (
	"github.com/zett-8/go-clean-echo/models"
	"github.com/zett-8/go-clean-echo/stores"
	utilss "github.com/zett-8/go-clean-echo/utils"
)

type (
	AccountService interface {
		GetTransactionByMatID(accountID int, paging *utilss.Paging) ([]*models.Transaction, error)
		GetAccountByMatID(mattermostID string) (*models.Account, error)
		GivePoint(trans *models.Transaction) error
		RejectPoint(trans *models.Transaction) error
		ExchangePointRequest(exRequest *models.ExchangeRequests) error
	}

	accountService struct {
		stores *stores.Stores
	}
)

func (s *accountService) GetTransactionByMatID(accountID int, paging *utilss.Paging) ([]*models.Transaction, error) {
	r, err := s.stores.Transaction.GetByAccountId(nil, accountID, paging)
	return r, err
}

func (s *accountService) GetAccountByMatID(mattermostID string) (*models.Account, error) {
	r, err := s.stores.Account.GetByMattId(nil, mattermostID)
	return r, err
}

func (s *accountService) GivePoint(trans *models.Transaction) error {
	return s.stores.Transaction.Save(nil, trans)
}

func (s *accountService) RejectPoint(trans *models.Transaction) error {
	return s.stores.Transaction.Save(nil, trans)
}

func (s *accountService) ExchangePointRequest(exRequest *models.ExchangeRequests) error {
	return s.stores.ExchangeRequest.Save(nil, exRequest)
}
