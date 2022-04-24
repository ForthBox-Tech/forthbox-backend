package model

import (
	"fmt"
	"forthboxbe/pkg/util"
	"time"

	"github.com/jinzhu/gorm"
)

const (
	TypeMobileVerify = "mobile_verify"
	TypeEmailVerify  = "email_verify"
	VtStatusSent     = 0
	VtStatusVerified = 1
)

type VerifyToken struct {
	ID        uint64    `gorm:"primaryKey;not null"`
	Token     string    `gorm:"unique;not null"`
	Code      string    `gorm:"not null;index;default:''"`
	Type      string    `gorm:"not null;default:''"`
	Data      string    `gorm:"not null;index;default:''"`
	Status    int       `gorm:"not null;default:0"`
	ExpiredAt time.Time `gorm:"not null;index"`
	SourceIP  string    `gorm:"not null;default:''"`
	VerifyIP  string    `gorm:"not null;default:''"`
	CreatedAt time.Time `gorm:"not null;index"`
	UpdatedAt time.Time `gorm:"not null;index"`
}
