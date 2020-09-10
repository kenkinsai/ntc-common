package model

type District struct {
	DistrictID   int    `xorm:"id" json:"district_id"`
	DistrictName string `xorm:"name" json:"district_name"`
	DistrictType string `xorm:"type" json:"district_type"`
	ProvinceID   int    `xorm:"province_id" json:"province_id"`
}
