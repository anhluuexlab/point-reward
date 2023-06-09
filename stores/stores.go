package stores

import (
	"gorm.io/gorm"
)

type Stores struct {
	DB              *gorm.DB
	Account         AccountStore
	Transaction     TransactionStore
	ExchangeRequest ExchangeRequestStore
}

func New(db *gorm.DB) *Stores {
	return &Stores{
		DB:              db,
		Account:         &accountStore{db},
		Transaction:     &transactionStore{db},
		ExchangeRequest: &exchangeRequestStore{db},
	}
}
