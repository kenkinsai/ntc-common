package model

import "time"

type Company struct {
	CompanyID             int       `xorm:"id" json:"company_id"`
	CompanyName           string    `xorm:"company_name" json:"company_name"`
	CompanySize           int       `xorm:"size" json:"size"`
	CompanyAddress        string    `xorm:"address" json:"address"`
	TaxCode               string    `xorm:"tax_code" json:"tax_code"`
	CompanyPhone          string    `xorm:"company_phone" json:"company_phone"`
	CompanyEmail          string    `xorm:"company_email" json:"company_email"`
	CompanyRepresentative string    `xorm:"representative" json:"representative"`
	Visible               bool      `xorm:"visible" json:"visible"`
	CreatedAt             time.Time `xorm:"created_at" json:"created_at"`
	UpdatedAt             time.Time `xorm:"updated_at" json:"updated_at"`
}
