package model

type UserContact struct {
	ContactID       int    `xorm:"contact_id" json:"contact_id"`
	ContactBy       int    `xorm:"contact_by" json:"contact_by"`
	IsActivated     int    `xorm:"is_activated" json:"is_activated"`
	IsDeleted       int    `xorm:"is_deleted" json:"is_deleted"`
	IsBlocked       int    `xorm:"is_blocked" json:"is_blocked"`
	EmailWork       string `xorm:"email_work" json:"email_work"`
	EmailHome       string `xorm:"email_home" json:"email_home"`
	PhoneNumberWork string `xorm:"phone_work" json:"phone_work"`
	PhoneNumberHome string `xorm:"phone_home" json:"phone_home"`
	CompanyName     string `xorm:"company_name" json:"company_name,omitempty"`
	Note            string `xorm:"note" json:"note,omitempty"`
	FirstName       string `xorm:"first_name" json:"first_name,omitempty"`
	LastName        string `xorm:"last_name" json:"last_name,omitempty"`
	VersionNo       string `xorm:"version_no" json:"version_no,omitempty"`
	CreatedAt       string `xorm:"created_at" json:"created_at,omitempty"`
	UpdatedAt       string `xorm:"updated_at" json:"updated_at,omitempty"`
	CreatedBy       int    `xorm:"created_by" json:"created_by,omitempty"`
	UpdatedBy       int    `xorm:"updated_by" json:"updated_by,omitempty"`
}

// IsExists struct
func (m UserContact) IsExists() (ok bool) {
	if m.ContactID != 0 {
		ok = true
	}
	return
}
