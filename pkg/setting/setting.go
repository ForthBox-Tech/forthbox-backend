// Package setting centralizes config loading for the backend process.
package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	IsDev       bool
	DebugLevel  int
	LogSavePath string
	LogSaveName string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	UserDsn string
}

var DatabaseSetting = &Database{}

type Mail struct {
	Domain string
	MGkey  string
}

var MailSetting = &Mail{}

var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("mail", MailSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

func IsDev() bool {
	return AppSetting.IsDev
}

