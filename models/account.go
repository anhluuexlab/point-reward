package models

import "time"

type Account struct {
	ID             int        `gorm:"column:id"  json:"-"`
	MattermostID   string     `gorm:"column:mattermost_id"  json:"mattermost_id"`
	UserName       string     `gorm:"column:username"  json:"username"`
	Email          string     `gorm:"column:email"  json:"email"`
	Role           string     `gorm:"column:role"  json:"role"`
	BalanceGranted int        `gorm:"column:balance_granted"  json:"balance_granted"`
	BalanceEarned  int        `gorm:"column:balance_earned"  json:"balance_earned"`
	CreatedAt      *time.Time `gorm:"column:created_at"  json:"-"`
	UpdatedAt      *time.Time `gorm:"column:updated_at"  json:"-"`
}
