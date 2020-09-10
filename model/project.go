package model

import "time"

type Project struct {
	ProjectID    int       `xorm:"id" json:"id"`
	Name         string    `xorm:"name" json:"name"`
	DistrictID   int       `xorm:"district_id" json:"district_id"`
	ProvinceID   int       `xorm:"province_id" json:"province_id"`
	Lat          string    `xorm:"lat" json:"lat"`
	Long         string    `xorm:"long" json:"long"`
	Owner        string    `xorm:"owner" json:"owner"`
	Address      string    `xorm:"address" json:"address"`
	AreaApart    string    `xorm:"area_apart" json:"area_apart"`
	TimeCons     string    `xorm:"time_cons" json:"time_cons"`
	HandoverTime string    `xorm:"handover_time" json:"handover_time"`
	PolicyLegal  string    `xorm:"policy_legal" json:"policy_legal"`
	BlockNumber  int       `xorm:"block_number" json:"block_number"`
	FloorNumber  int       `xorm:"floor_number" json:"floor_number"`
	ApartNumber  int       `xorm:"apart_number" json:"apart_number"`
	Utilities    []string  `xorm:"utilities" json:"utilities"`
	Price        string    `xorm:"price" json:"price"`
	Images       []string  `xorm:"images" json:"images"`
	Visible      string    `xorm:"visible" json:"visible"`
	CreatedAt    time.Time `xorm:"created_at" json:"created_at"`
	CreatedBy    string    `xorm:"created_by" json:"created_by"`
	UpdatedAt    time.Time `xorm:"updated_by" json:"updated_by"`
	UpdatedBy    string    `xorm:"updated_at" json:"updated_at"`
}

type ProjectWithDistance struct {
	ProjectID    int       `xorm:"id" json:"id"`
	Name         string    `xorm:"name" json:"name"`
	DistrictID   int       `xorm:"district_id" json:"district_id"`
	ProvinceID   int       `xorm:"province_id" json:"province_id"`
	Lat          string    `xorm:"lat" json:"lat"`
	Long         string    `xorm:"long" json:"long"`
	Owner        string    `xorm:"owner" json:"owner"`
	Address      string    `xorm:"address" json:"address"`
	AreaApart    string    `xorm:"area_apart" json:"area_apart"`
	TimeCons     string    `xorm:"time_cons" json:"time_cons"`
	HandoverTime string    `xorm:"handover_time" json:"handover_time"`
	BlockNumber  int       `xorm:"block_number" json:"block_number"`
	PolicyLegal  string    `xorm:"policy_legal" json:"policy_legal"`
	FloorNumber  int       `xorm:"floor_number" json:"floor_number"`
	ApartNumber  int       `xorm:"apart_number" json:"apart_number"`
	Utilities    []string  `xorm:"utilities" json:"utilities"`
	Price        string    `xorm:"price" json:"price"`
	Images       []string  `xorm:"images" json:"images"`
	Visible      string    `xorm:"visible" json:"visible"`
	CreatedAt    time.Time `xorm:"created_at" json:"created_at"`
	CreatedBy    string    `xorm:"created_by" json:"created_by"`
	UpdatedAt    time.Time `xorm:"updated_by" json:"updated_by"`
	UpdatedBy    string    `xorm:"updated_at" json:"updated_at"`
	Distance     float64   `xorm:"distance" json:"distance"`
}
