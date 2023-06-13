package models

import "time"

type Transaction struct {
	ID         int        `gorm:"column:id;primary_key"  json:"id"`
	Action     string     `gorm:"column:action"  json:"action"`
	Amount     int        `gorm:"column:amount"  json:"amount"`
	SenderID   int        `gorm:"column:sender_id"  json:"-"`
	Sender     Account    `gorm:"column:sender_id"  json:"sender"`
	ReceiverID int        `gorm:"column:receiver_id"  json:"-"`
	Receiver   Account    `gorm:"column:receiver_id"  json:"receiver"`
	OperatorID int        `gorm:"column:operator_id"  json:"-"`
	Operator   *Account   `gorm:"column:operator_id"  json:"operator,omitempty"`
	CreatedAt  *time.Time `gorm:"column:created_at"  json:"-"`
	UpdatedAt  *time.Time `gorm:"column:updated_at"  json:"-"`
}

type GivePointForm struct {
	Amount     int    `json:"amount" validate:"required"`
	SenderID   string `json:"sender_id" validate:"required"`
	ReceiverID string `json:"receiver_id" validate:"required"`
}

type RejectPointForm struct {
	TransactionID int    `json:"transaction_id" validate:"required"`
	OperatorID    string `json:"operator_id" validate:"required"`
}
