package stores

import (
	"database/sql"

	"gorm.io/gorm"

	"github.com/zett-8/go-clean-echo/models"
)

type (
	ExchangeRequestStore interface {
		ExchangeRequest(tx *sql.Tx, requester *models.Account, trans *models.ExchangeRequests) error
	}

	exchangeRequestStore struct {
		db *gorm.DB
	}
)

func (s *exchangeRequestStore) ExchangeRequest(tx *sql.Tx, requester *models.Account, exRequest *models.ExchangeRequests) error {
	ts := s.db.Begin()
	err := ts.Error
	if err != nil {
		return err
	}
	// requester
	requester.BalanceEarned = requester.BalanceEarned - exRequest.Amount
	err = ts.Save(&requester).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// ex transaction
	err = ts.Save(&exRequest).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// transaction
	trans := &models.Transaction{
		Action:     "exchange",
		Amount:     exRequest.Amount,
		SenderID:   requester.ID,
		ReceiverID: 0,
	}
	err = ts.Save(&trans).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = ts.Commit().Error
	return err
}
