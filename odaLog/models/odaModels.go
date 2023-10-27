package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username string `binding:"required"`
}
type Borcs struct {
	gorm.Model
	Borc   float32
	Borclu string
	Uid    uint
	Users  Users `gorm:"foreignKey:Uid"`
}
