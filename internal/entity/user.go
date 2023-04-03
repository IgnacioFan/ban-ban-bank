package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID   uint64 `json:"id" gorm:"primarykey"`
	Name string `json:"name" gorm:"type:varchar(255)"`
}
