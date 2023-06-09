package models

import "time"

type Transaction struct {
	ID         int        `db:"id;primary_key"  json:"id"`
	Action     string     `db:"action"  json:"action"`
	Amount     int        `db:"amount"  json:"amount"`
	SenderID   int        `db:"sender_id"  json:"-"`
	Sender     Account    `db:"sender_id"  json:"sender"`
	ReceiverID int        `db:"receiver_id"  json:"-"`
	Receiver   Account    `db:"receiver_id"  json:"receiver"`
	OperatorID int        `db:"operator_id"  json:"-"`
	Operator   *Account   `db:"operator_id, omitempty"  json:"operator"`
	CreatedAt  *time.Time `db:"created_at, omitempty" json:"-" `
	UpdatedAt  *time.Time `db:"updated_at, omitempty" json:"-" `
}

type GivePointForm struct {
	Amount     int `json:"amount" validate:"required"`
	SenderID   int `json:"sender_id" validate:"required"`
	ReceiverID int `json:"receiver_id" validate:"required"`
}

type RejectPointForm struct {
	TransactionID int `json:"transaction_id" validate:"required"`
	OperatorID    int `json:"operator_id" validate:"required"`
}
