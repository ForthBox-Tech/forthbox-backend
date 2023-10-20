package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	TypeBannerPic         = "banner"
	TypeGameCenterListPic = "game_center"
)

type Pic struct {
	ID        uint64    `gorm:"primaryKey;not null"`
	Type      string    `gorm:"not null;index;default:''"`
	Title     string    `gorm:"not null;default:''"`
	Href      string    `gorm:"not null;default:''"`
	Src       string    `gorm:"not null;default:''"`
	Status    int       `gorm:"not null;default:1"`
	Creator   string    `gorm:"not null;default:''"`
	CreatorIP string    `gorm:"not null;default:''"`
	Sort      int       `gorm:"not null;index;default:1"`
	CreatedAt time.Time `gorm:"not null;index"`
	UpdatedAt time.Time `gorm:"not null;index"`
}

func GetPicModel() *gorm.DB {
	return UserDb.Model(&Pic{})
}

