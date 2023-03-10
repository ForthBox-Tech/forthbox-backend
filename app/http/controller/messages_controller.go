package controller

import (
	"fmt"
	"forthboxbe/app/service"
	"forthboxbe/pkg/setting"

	"github.com/gin-gonic/gin"
)

func SendEmailVerifyCode(c *gin.Context) {
	resp := NewResp(c)

	email := c.PostForm("email")
	if email == "" {
		resp.Err("params email is required")
		return
	}
	src_ip := c.ClientIP()
	vk, err := service.AddEmailVerifyToken(email, src_ip)
	if err != nil {
		resp.Err("发送失败")
		return
	}
	o := gin.H{"is_success": true}
	if setting.IsDev() {
		o["code"] = vk.Code
	}
	resp.Out(o)
}

func SendMobileVerifyCode(c *gin.Context) {
	resp := NewResp(c)

	mobile := c.PostForm("mobile")
	if mobile == "" {
		resp.Err("params mobile is required")
		return
	}
	m_rigion := c.DefaultPostForm("m_rigion", "86")
	src_ip := c.ClientIP()
	mobile_str := fmt.Sprintf("%s.%s", m_rigion, mobile)
	vk, err := service.AddMobileVerifyToken(mobile_str, src_ip)
	if err != nil {
		resp.Err("发送失败")
		return
	}
	o := gin.H{"is_success": true}
	if setting.IsDev() {
		o["code"] = vk.Code
	}
	resp.Out(o)
}


