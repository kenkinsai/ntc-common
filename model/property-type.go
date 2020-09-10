package model

type PropertyType struct {
	ID      int    `xorm:"id" json:"id"`
	Slug    string `xorm:"slug" json:"slug"`
	Name    string `xorm:"name" json:"name"`
	Visible bool   `xorm:"visible" json:"visible"`
}
