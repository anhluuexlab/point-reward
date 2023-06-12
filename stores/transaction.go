package stores

import (
	"database/sql"

	"gorm.io/gorm"

	"github.com/zett-8/go-clean-echo/models"
	"github.com/zett-8/go-clean-echo/utils"
)

type (
	TransactionStore interface {
		Save(tx *sql.Tx, trans *models.Transaction) error
		GivePoint(tx *sql.Tx, sender *models.Account, receiver *models.Account, trans *models.Transaction) error
		GetByAccountId(tx *sql.Tx, accountID int, paging *utils.Paging) ([]*models.Transaction, error)
		GetTransactionByID(tx *sql.Tx, ID int) (*models.Transaction, error)
		RejectPoint(tx *sql.Tx, sender *models.Account, receiver *models.Account, trans *models.Transaction) error
	}

	transactionStore struct {
		db *gorm.DB
	}
)

func (s *transactionStore) Save(tx *sql.Tx, trans *models.Transaction) error {
	return s.db.Save(&trans).Error
}

func (s *transactionStore) GivePoint(tx *sql.Tx, sender *models.Account, receiver *models.Account, trans *models.Transaction) error {
	ts := s.db.Begin()
	err := ts.Error
	if err != nil {
		return err
	}
	// sender
	if sender.BalanceGranted >= trans.Amount {
		sender.BalanceGranted = sender.BalanceGranted - trans.Amount
	} else {
		sender.BalanceEarned = sender.BalanceEarned - trans.Amount
	}
	err = ts.Save(&sender).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// receiver
	receiver.BalanceEarned = receiver.BalanceEarned + trans.Amount
	err = ts.Save(&receiver).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// transaction
	err = ts.Save(&trans).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = ts.Commit().Error
	return err
}

func (s *transactionStore) GetByAccountId(tx *sql.Tx, accountID int, paging *utils.Paging) ([]*models.Transaction, error) {
	trans := []*models.Transaction{}

	err := s.db.Table("transactions").
		Preload("Receiver").
		Where("sender_id = ? or receiver_id = ?", accountID, accountID).
		Find(&trans).Error

	return trans, err
}

func (s *transactionStore) GetTransactionByID(tx *sql.Tx, ID int) (*models.Transaction, error) {
	trans := &models.Transaction{}

	err := s.db.Table("transactions").
		Where("id = ?", ID).
		First(&trans).Error

	return trans, err
}

func (s *transactionStore) RejectPoint(tx *sql.Tx, sender *models.Account, receiver *models.Account, trans *models.Transaction) error {
	ts := s.db.Begin()
	err := ts.Error
	if err != nil {
		return err
	}
	// sender
	if sender.BalanceEarned >= trans.Amount {
		sender.BalanceEarned = sender.BalanceEarned - trans.Amount
	} else {
		sender.BalanceGranted = sender.BalanceGranted - trans.Amount
	}
	err = ts.Save(&sender).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// receiver
	receiver.BalanceGranted = receiver.BalanceGranted + trans.Amount
	err = ts.Save(&receiver).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// transaction
	err = ts.Save(&trans).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = ts.Commit().Error
	return err
}
