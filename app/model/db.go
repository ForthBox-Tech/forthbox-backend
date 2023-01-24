package model

import (
	"fmt"
	"forthboxbe/pkg/setting"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var UserDb *gorm.DB

func init() {
	var err error
	db, err := gorm.Open("mysql", setting.DatabaseSetting.UserDsn)

