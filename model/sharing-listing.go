package model

import "time"

type SharingListing struct {
	ID                int       `xorm:"id" json:"id"`
	ListingID         int       `xorm:"listing_id" json:"listing_id"`
	UserID            int       `xorm:"user_id" json:"user_id"`
	Status            string    `xorm:"status" json:"status"`
	Note              string    `xorm:"note" json:"note"`
	SharingPercentage int       `xorm:"sharing_percentage" json:"sharing_percentage"`
	CreatedAt         time.Time `xorm:"created_at" json:"created_at"`
	CreatedBy         int       `xorm:"created_by" json:"created_by"`
	UpdatedAt         time.Time `xorm:"updated_at" json:"updated_at"`
	UpdatedBy         int       `xorm:"updated_by" json:"updated_by"`
}
