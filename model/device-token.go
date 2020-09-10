package model

import "time"

type DeviceToken struct {
	ID        int       `xorm:"id" json:"id"`
	UserID    int       `xorm:"user_id" json:"user_id"`
	App       string    `xorm:"app" json:"app"`
	Source    string    `xorm:"source" json:"source"`
	Token     string    `xorm:"noti_token" json:"noti_token"`
	Visible   bool      `xorm:"visible" json:"visible"`
	CreatedAt time.Time `xorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at" json:"updated_at"`
}
