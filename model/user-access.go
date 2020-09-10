package model

type UserAccess struct {
	UserID         int    `xorm:"user_id" json:"user_id"`
	IsActived      int    `xorm:"is_actived" json:"is_actived"`
	IsDeleted      int    `xorm:"is_deleted" json:"is_deleted"`
	IsBlocked      int    `xorm:"is_blocked" json:"is_blocked"`
	Email          string `xorm:"email" json:"email,omitempty"`
	PhoneNumber    string `xorm:"phone_number" json:"phone_number,omitempty"`
	Password       string `xorm:"password" json:"password,omitempty"`
	PasswordSalt   string `xorm:"password_salt" json:"password_salt,omitempty"`
	VersionNo      string `xorm:"version_no" json:"version_no,omitempty"`
	CreatedAt      string `xorm:"created_at" json:"created_at,omitempty"`
	UpdatedBy      int    `xorm:"updated_by" json:"updated_by,omitempty"`
}

// IsExists struct
func (m UserAccess) IsExists() (ok bool) {
	if m.UserID != 0 {
		ok = true
	}
	return
}
