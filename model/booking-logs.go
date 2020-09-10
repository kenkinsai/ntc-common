package model

import "time"

type BookingLogs struct {
	BookingID     int       `xorm:"booking_id" json:"booking_id"`
	BookingStatus string    `xorm:"booking_status" json:"booking_status"`
	Action        string    `xorm:"action" json:"action"`
	Note          string    `xorm:"note" json:"note"`
	BookingTime   int       `xorm:"booking_time" json:"booking_time"`
	Visible       bool      `xorm:"visible" json:"visible"`
	CreatedBy     int       `xorm:"created_by" json:"created_by"`
	CreatedAt     time.Time `xorm:"created_at" json:"created_at"`
}
