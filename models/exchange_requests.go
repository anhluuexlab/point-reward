package models

import "time"

type ExchangeRequests struct {
	ID          int        `db:"id"  json:"-"`
	Status      string     `db:"status"  json:"status"`
	Amount      int        `db:"amount"  json:"amount"`
	RequesterID int        `db:"requester_id"  json:"-"`
	Requester   Account    `db:"requester_id"  json:"requester"`
	OperatorID  int        `db:"operator_id"  json:"-"`
	Operator    *Account   `db:"operator_id, omitempty"  json:"operator"`
	CreatedAt   *time.Time `db:"created_at, omitempty" json:"-" `
	UpdatedAt   *time.Time `db:"updated_at, omitempty" json:"-" `
}

type ExchangeRequestForm struct {
	Amount      int `json:"amount" validate:"required"`
	RequesterID int `json:"requester_id" validate:"required"`
}
