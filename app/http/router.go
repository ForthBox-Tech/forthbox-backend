package router

import (
	"forthboxbe/app/http/controller"
	"forthboxbe/pkg/setting"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	// Allow known origins while keeping local development flexible.
	// cors
	corsConfig := cors.DefaultConfig()
	if setting.IsDev() {
		corsConfig.AllowAllOrigins = true
	} else {
		corsConfig.AllowOrigins = []string{
			"https://forthbox.io", "https://www.forthbox.io",
			"http://forthbox.io", "http://www.forthbox.io",
			"https://forthbox.com", "https://www.forthbox.com",
			"http://forthbox.com", "http://www.forthbox.com",
		}
	}
	r.Use(cors.New(corsConfig))

	// placeholder
	r.GET("/", controller.HomeIndex)

	r.POST("/users/login", controller.UserLogin)
	r.POST("/users/signup_by_email", controller.SignUpByEmail)
	r.POST("/users/signup_by_mobile", controller.SignUpByMobile)
	r.POST("/users/reset_password", controller.ResetPassword)
	r.POST("/users/set_password", controller.SetPassword)
	r.GET("/users/check_exist", controller.CheckUserExist)
	r.GET("/users/get_auth_info", controller.GetAuthInfo)

	r.POST("/messages/send_email_verify_code", controller.SendEmailVerifyCode)
	r.POST("/messages/send_mobile_verify_code", controller.SendMobileVerifyCode)

	return r
}


