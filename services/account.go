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
		GivePoint(sender *models.Account, receiver *models.Account, trans *models.Transaction) error
		RejectPoint(sender *models.Account, receiver *models.Account, trans *models.Transaction) error
		ExchangePointRequest(requester *models.Account, exRequest *models.ExchangeRequests) error
		GetTransactionByID(ID int) (*models.Transaction, error)
		GetAccountByID(ID int) (*models.Account, error)
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

func (s *accountService) GetAccountByID(ID int) (*models.Account, error) {
	r, err := s.stores.Account.GetByID(nil, ID)
	return r, err
}

func (s *accountService) GivePoint(sender *models.Account, receiver *models.Account, trans *models.Transaction) error {
	return s.stores.Transaction.GivePoint(nil, sender, receiver, trans)
}

func (s *accountService) RejectPoint(sender *models.Account, receiver *models.Account, trans *models.Transaction) error {
	return s.stores.Transaction.RejectPoint(nil, sender, receiver, trans)
}

func (s *accountService) ExchangePointRequest(requester *models.Account, exRequest *models.ExchangeRequests) error {
	return s.stores.ExchangeRequest.ExchangeRequest(nil, requester, exRequest)
}

func (s *accountService) GetTransactionByID(ID int) (*models.Transaction, error) {
	r, err := s.stores.Transaction.GetTransactionByID(nil, ID)
	return r, err
}
