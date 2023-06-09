package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name" gorm:"size:50"`
	Age  int64  `json:"age"`
	Sex  string `json:"sex" gorm:"size:2"`
}
