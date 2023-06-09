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
		GetByAccountId(tx *sql.Tx, accountID int, paging *utils.Paging) ([]*models.Transaction, error)
	}

	transactionStore struct {
		db *gorm.DB
	}
)

func (s *transactionStore) Save(tx *sql.Tx, trans *models.Transaction) error {
	return s.db.Save(&trans).Error
}

func (s *transactionStore) GetByAccountId(tx *sql.Tx, accountID int, paging *utils.Paging) ([]*models.Transaction, error) {
	trans := []*models.Transaction{}

	err := s.db.Table("transactions").
		Preload("Receiver").
		Where("sender_id = ? or receiver_id = ?", accountID, accountID).
		Find(&trans).Error

	return trans, err
}
