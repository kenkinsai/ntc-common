package model

type Provinces struct {
	ProvinceID   int    `xorm:"id" json:"province_id"`
	ProvinceName string `xorm:"name" json:"province_name"`
	ProvinceType string `xorm:"type" json:"province_type"`
}
