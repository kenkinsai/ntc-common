package model

import "time"

type UserCompany struct {
	UserID    int       `xorm:"user_id" json:"user_id"`
	CompanyID int       `xorm:"company_id" json:"company_id"`
	Visible   bool      `xorm:"visible" json:"visible"`
	CreatedAt time.Time `xorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at" json:"updated_at"`
}
