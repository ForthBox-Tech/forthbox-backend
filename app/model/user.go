package model

import (
	"forthboxbe/pkg/util"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID             uint64    `gorm:"primaryKey;not null"`
	UserName       string    `gorm:"unique;not null"`
	Passwd         string    `gorm:"not null;default:''"`
	Email          string    `gorm:"not null;index;default:''"`
	EmailIsVerify  int       `gorm:"not null;default:0"`
	MRigion        string    `gorm:"not null;default:86"`
	Mobile         string    `gorm:"not null;index;default:''"`
	MobileIsVerify int       `gorm:"not null;default:0"`
	Gender         string    `gorm:"not null;default:'male'"`
	Status         int       `gorm:"not null;default:0"`
	InviteCode     string    `gorm:"not null;default:'';comment:'个人注册码'"`
	UseInviteCode  string    `gorm:"not null;default:'';comment:'使用的注册码'"`
	RegisterIP     string    `gorm:"not null;default:''"`
	LastSignInAt   time.Time `gorm:"not null;autoCreateTime"`
	LastSignInIP   string    `gorm:"not null;default:''"`
	CreatedAt      time.Time `gorm:"not null;index"`
	UpdatedAt      time.Time `gorm:"not null;index"`
}

const UserPasswordSalt = "374sx&g6e#7b^$3"

const (
	UserStatusPendding = 0
	UserStatusNormal   = 1
)

func GetUserModel() *gorm.DB {
	return UserDb.Model(&User{})
}

func (u *User) IsUserNameExist(username string) (bool, error) {
	var count int
	if err := GetUserModel().Where(&User{UserName: username}).Count(&count).Error; err != nil {
		return true, err
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

// placeholder
func HashPassword(input string) string {
	hash := util.Sha256String(input + UserPasswordSalt)

	
	return hash
}

// placeholder
func (u *User) CheckPassword(input string) bool {
	hash := HashPassword(input)
	if hash == u.Passwd {
		return true
	} else {
		return false
	}
}



