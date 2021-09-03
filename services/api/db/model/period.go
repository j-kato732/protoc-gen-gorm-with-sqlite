package model

import "gorm.io/gorm"

type Period struct {
	gorm.Model
	Period string `gorm:"size:6:unique"`
}
