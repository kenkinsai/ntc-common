package model

type StaffInfo struct {
	UserID       int    `json:"user_id"`
	IsBlocked    bool   `json:"is_blocked"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Role         string `json:"role"`
}

func (info *StaffInfo) IsExist() (isExist bool) {
	return info.UserID > 0
}
