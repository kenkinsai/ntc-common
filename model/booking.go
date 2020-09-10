package model

import "time"

type Booking struct {
	BookingID      int       `xorm:"booking_id" json:"booking_id"`
	ItemID         int       `xorm:"item_id" json:"item_id"`
	BookingMaker   int       `xorm:"booking_maker" json:"booking_maker"`
	BookingHandler int       `xorm:"booking_handler" json:"booking_handler"`
	Status         string    `xorm:"status" json:"status"`
	IsDeleted      int       `xorm:"is_deleted" json:"is_deleted"`
	BookingType    int       `xorm:"booking_type" json:"booking_type"`
	VersionNo      string    `xorm:"version_no" json:"version_no"`
	CreatedBy      int       `xorm:"created_by" json:"created_by"`
	UpdatedBy      int       `xorm:"updated_by" json:"updated_by"`
	CreatedAt      time.Time `xorm:"created_at" json:"created_at"`
	UpdatedAt      time.Time `xorm:"updated_at" json:"updated_at"`
}
