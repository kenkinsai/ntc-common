package model

import "time"

type UserNotify struct {
	ID        int       `xorm:"id" json:"id"`
	UserID    int       `xorm:"user_id" json:"user_id"`
	App       string    `xorm:"app" json:"app"`
	Title     string    `xorm:"title" json:"title"`
	Body      string    `xorm:"body" json:"body"`
	Payload   string    `xorm:"payload" json:"payload"`
	Icon      string    `xorm:"icon" json:"icon"`
	IsOpen    bool      `xorm:"is_open" json:"is_open"`
	Visible   bool      `xorm:"visible" json:"visible"`
	CreatedAt time.Time `xorm:"created_at" json:"created_at"`
}
