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

func GetVerifyTokenModel() *gorm.DB {
	return UserDb.Model(&VerifyToken{})
}

// placeholder
func (vt *VerifyToken) CreateVerifyToken(vtype string, vdata string, code string, src_ip string) (uint64, error) {
	token_str := util.Sha256String(fmt.Sprintf("%s:%s:%s:%s", vtype, vdata, src_ip, util.RandString(20)))

	vk := VerifyToken{
		Token: token_str, Code: code, Type: vtype, Data: vdata,
		Status: 0, SourceIP: src_ip,
	}
	vk.ExpiredAt = time.Now().Local().Add(time.Minute * time.Duration(30))
	if err := GetVerifyTokenModel().Create(&vk).Error; err != nil {
		return 0, err
	}

	return uint64(vk.ID), nil
}



