package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string   `gorm:"unique" json:"username"`
	Role     string   `json:"role"`
	NikID    int64    `json:"nikid"`
	NIK      Employee `gorm:"foreignKey:NikID;references:Nik"`
	Password string   `json:"password"`
}
