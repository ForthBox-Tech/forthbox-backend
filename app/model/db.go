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

	if err != nil {
		log.Fatalf("Db init err: %v", err)
	} else {
		fmt.Println("UserDb connected")
		if setting.AppSetting.DebugLevel > 0 {
			UserDb = db.Debug()
		} else {
			UserDb = db
		}
	}
}


