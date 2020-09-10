package model

import "time"

type Image struct {
	ImageID   int       `xorm:"id" json:"id"`
	ListingID int       `xorm:"listing_id" json:"listing_id"`
	Name      string    `xorm:"photo" json:"photo"`
	Link      string    `xorm:"link" json:"link"`
	Label     string    `xorm:"label" json:"label"`
	Visible   int       `xorm:"visible" json:"visible"`
	UploadBy  int       `xorm:"upload_by" json:"upload_by"`
	UploadAt  time.Time `xorm:"upload_at" json:"upload_at"`
}
