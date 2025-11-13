package entity

import (
	"time"
	"user-management/pkg/enums"
)

type User struct {
	ID              int64         `gorm:"primaryKey;autoIncrement"`
	AuthID          string        `gorm:"type:varchar(36);uniqueIndex;not null"`
	Username        string        `gorm:"type:varchar(100);uniqueIndex;not null"`
	FirstName       *string       `gorm:"type:varchar(100)"`
	LastName        *string       `gorm:"type:varchar(100)"`
	Email           *string       `gorm:"type:varchar(100);uniqueIndex"`
	Phone           *string       `gorm:"type:varchar(20);index"`
	BirthDate       *time.Time    `gorm:"type:date"`
	Gender          *enums.GENDER `gorm:"type:varchar(10)"`
	AvatarURL       *string       `gorm:"type:text"`
	Bio             *string       `gorm:"type:varchar(500)"`
	Description     *string       `gorm:"type:text"`
	EmailVerifiedAt *time.Time    `gorm:"type:timestamp"`
	PhoneVerifiedAt *time.Time    `gorm:"type:timestamp"`
	CreatedAt       time.Time     `gorm:"autoCreateTime"`
	UpdatedAt       time.Time     `gorm:"autoUpdateTime"`
}
