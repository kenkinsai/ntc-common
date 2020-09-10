package model

type Career struct {
	CareerID   int    `xorm:"career_id" json:"career_id"`
	CareerName string `xorm:"name" json:"name"`
	Priority   int    `xorm:"priority" json:"priority"`
	ParentID   int    `xorm:"parent_id" json:"parent_id"`
}

// IsExists struct
func (m Career) IsExists() (ok bool) {
	if m.CareerID != 0 {
		ok = true
	}
	return
}
