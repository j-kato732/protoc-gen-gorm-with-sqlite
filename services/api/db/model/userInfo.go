package model

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	UserId        int32  `gorm:"not null"`
	LastName      string `gorm:"size:10;not null"`
	FirstName     string `gorm:"size:10;not null"`
	Period        string `gorm:"size:6;not null`
	DepartmentId  int32  `gorm:"not null"`
	JobId         int32  `gorm:"not null"`
	EnrollmentFlg bool
	AdminFlg      bool
}
