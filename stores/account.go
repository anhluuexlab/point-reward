package stores

import (
	"database/sql"

	"gorm.io/gorm"

	"github.com/zett-8/go-clean-echo/models"
)

type (
	AccountStore interface {
		GetAll(tx *sql.Tx) ([]*models.Account, error)
		GetByMattId(tx *sql.Tx, mattermostID string) (*models.Account, error)
	}

	accountStore struct {
		db *gorm.DB
	}
)

func (s *accountStore) GetAll(tx *sql.Tx) ([]*models.Account, error) {
	accounts := []*models.Account{}

	err := s.db.Model(models.Account{}).Find(&accounts).Error

	return accounts, err
}

func (s *accountStore) GetByMattId(tx *sql.Tx, mattermostID string) (*models.Account, error) {
	acc := &models.Account{}

	err := s.db.Table("accounts").Where("mattermost_id = ?", mattermostID).First(acc).Error

	return acc, err
}
