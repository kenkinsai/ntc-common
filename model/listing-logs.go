package model

import "time"

type ListingLogs struct {
	ListingID int       `xorm:"listing_id" json:"listing_id"`
	Status    string    `xorm:"status" json:"status"`
	Note      string    `xorm:"note" json:"note"`
	Visible   bool      `xorm:"visible" json:"visible"`
	CreatedBy int       `xorm:"created_by" json:"created_by"`
	CreatedAt time.Time `xorm:"created_at" json:"created_at"`
}
