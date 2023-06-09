package stores

import (
	"database/sql"

	"gorm.io/gorm"

	"github.com/zett-8/go-clean-echo/models"
)

type (
	ExchangeRequestStore interface {
		Save(tx *sql.Tx, trans *models.ExchangeRequests) error
	}

	exchangeRequestStore struct {
		db *gorm.DB
	}
)

func (s *exchangeRequestStore) Save(tx *sql.Tx, exRequest *models.ExchangeRequests) error {
	return s.db.Save(&exRequest).Error
}
