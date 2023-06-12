package models

import "time"

type ExchangeRequests struct {
	ID          int        `gorm:"column:id"  json:"-"`
	Status      string     `gorm:"column:status"  json:"status"`
	Amount      int        `gorm:"column:amount"  json:"amount"`
	RequesterID int        `gorm:"column:requester_id"  json:"-"`
	Requester   Account    `gorm:"column:requester_id"  json:"requester"`
	OperatorID  int        `gorm:"column:operator_id"  json:"-"`
	Operator    *Account   `gorm:"column:operator_id"  json:"operator"`
	CreatedAt   *time.Time `gorm:"column:created_at"  json:"-"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"  json:"-"`
}

type ExchangeRequestForm struct {
	Amount      int    `json:"amount" validate:"required"`
	RequesterID string `json:"requester_id" validate:"required"`
}
