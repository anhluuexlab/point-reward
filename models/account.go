package models

import "time"

type Account struct {
	ID             int        `db:"id"  json:"-"`
	MattermostID   string     `db:"mattermost_id"  json:"mattermost_id"`
	UserName       string     `db:"username"  json:"username"`
	Email          string     `db:"email"  json:"email"`
	Role           string     `db:"role"  json:"role"`
	BalanceGranted int        `db:"balance_granted"  json:"balance_granted"`
	BalanceEarned  int        `db:"balance_earned"  json:"balance_earned"`
	CreatedAt      *time.Time `db:"created_at, omitempty" json:"-" `
	UpdatedAt      *time.Time `db:"updated_at, omitempty" json:"-" `
}
