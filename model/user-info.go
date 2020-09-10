package model

import "time"

type ShowUserInfo struct {
	ID             int    `xorm:"id" json:"id"`
	UserID         int    `xorm:"user_id" json:"user_id"`
	Gender         int    `xorm:"gender,omitempty" json:"gender"`
	Email          string `xorm:"email" json:"email"`
	PhoneNumber    string `xorm:"phone_number" json:"phone_number"`
	DisplayName    string `xorm:"display_name" json:"display_name"`
	DisplayAddress string `xorm:"display_address" json:"display_address"`
	FireBaseUserID string `xorm:"firebase_user_id" json:"firebase_user_id,omitempty"`
	Avatar         string `xorm:"avatar_link" json:"avatar_link"`
	Status         int    `xorm:"status" json:"status"`
	Verified       bool   `xorm:"verified" json:"verified"`
}

type UserInfo struct {
	UserID             int       `xorm:"user_id" json:"user_id"`
	Gender             int       `xorm:"gender" json:"gender"`
	Email              string    `xorm:"email" json:"email"`
	PhoneNumber        string    `xorm:"phone_number" json:"phone_number"`
	DisplayName        string    `xorm:"display_name" json:"display_name"`
	DisplayAddress     string    `xorm:"display_address" json:"display_address"`
	YearOfBirth        int       `xorm:"year_of_birth" json:"year_of_birth"`
	CareerId           int       `xorm:"career_id" json:"career_id"`
	AddressID          int       `xorm:"address_id" json:"address_id"`
	FirstName          string    `xorm:"first_name" json:"first_name"`
	LastName           string    `xorm:"last_name" json:"last_name"`
	Avatar             string    `xorm:"avatar_link" json:"avatar_link"`
	PersonalCardAhead  string    `xorm:"personal_card_ahead" json:"personal_card_ahead"`
	PersonalCardBehind string    `xorm:"personal_card_behind" json:"personal_card_behind"`
	DateOfBirth        time.Time `xorm:"date_of_birth" json:"date_of_birth,omitempty"`
	FireBaseToken      string    `xorm:"firebase_token" json:"firebase_token,omitempty"`
	FireBaseUserID     string    `xorm:"firebase_user_id" json:"firebase_user_id,omitempty"`
	PersonalID         string    `xorm:"personal_id" json:"personal_id"`
	Relationship       string    `xorm:"relationship" json:"relationship"`
	Language           string    `xorm:"language" json:"language"`
	VersionNo          string    `xorm:"version_no" json:"version_no"`
	CreatedAt          string    `xorm:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt          string    `xorm:"updated_at,omitempty" json:"updated_at,omitempty"`
	UpdatedBy          string    `xorm:"updated_by,omitempty" json:"updated_by,omitempty"`
	Status             int       `xorm:"status" json:"status"`
	Verified           bool      `xorm:"verified" json:"verified"`
}

// IsExists struct
func (m UserInfo) IsExists() (ok bool) {
	if m.UserID != 0 {
		ok = true
	}
	return
}
