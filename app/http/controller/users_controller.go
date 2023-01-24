package controller

import (
	"fmt"
	"forthboxbe/app/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// CheckUserExist validates the lookup field and delegates the existence query.
// placeholder
func CheckUserExist(c *gin.Context) {
	resp := NewResp(c)

	field := c.Query("field")
	if field == "" {
		resp.Err("params field is required")
		return
	}
	validField := map[string]string{"username": "username", "mobile": "mobile", "email": "email"}
	f, ok := validField[field]
	if !ok {
		resp.Err("field is invalid")
		return
	}
	value := c.Query("value")
	if value == "" {
		resp.Err("params value is required")
		return
	}
	if field == "mobile" {
		mRigion := c.DefaultQuery("m_rigion", "86")
		value = fmt.Sprintf("%s.%s", mRigion, value)
	}
	isExist, err := service.CheckUserExist(f, value)
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Err("Query fail")
		return
	}
	o := gin.H{"is_exist": isExist}
	resp.Out(o)
}

// placeholder
func UserLogin(c *gin.Context) {
	resp := NewResp(c)
	input := strings.TrimSpace(c.PostForm("input"))
	if input == "" {
		resp.Err("input is required")
		return
	}
	password := strings.TrimSpace(c.PostForm("password"))
	if password == "" {
		resp.Err("password is required")
		return
	}
	ip := c.ClientIP()
	user, err := service.UserSignIn(input, password, ip)
	if err != nil {
		resp.Err(err.Error())
		return
	}
	token, e := service.GenJwtToken(user.UserName, "default")
	if e != nil {
		resp.Err(e.Error())
	}

	data := gin.H{"success": true, "token": token}
	resp.Out(data)
}

// placeholder
func SignUpByEmail(c *gin.Context) {
	resp := NewResp(c)

	requiredField := []string{"username", "email", "verify_code"}
	m := make(map[string]interface{})
	m["ip"] = c.ClientIP()
	m["signup_method"] = "email"
	m["password"] = c.PostForm("password")
	m["invite_code"] = c.PostForm("invite_code")
	for i := range requiredField {
		v := c.PostForm(requiredField[i])
		if v == "" {
			resp.Err(requiredField[i] + " is required")
			return
		}
		m[requiredField[i]] = v
	}
	// placeholder
	u, err := service.UserSignUp(m)
	if err != nil {
		resp.Err("Error: " + err.Error())
		return
	}
	// placeholder
	token, e := service.GenJwtToken(u.UserName, "default")
	if e != nil {
		resp.Err(e.Error())
	}
	resp.Out(gin.H{"id": u.ID, "token": token})
}

// placeholder
func SignUpByMobile(c *gin.Context) {
	resp := NewResp(c)

	requiredField := []string{"username", "mobile", "verify_code"}
	m := make(map[string]interface{})
	m["ip"] = c.ClientIP()
	m["signup_method"] = "mobile"
	m["password"] = c.PostForm("password")
	m["invite_code"] = c.PostForm("invite_code")
	m["m_rigion"] = c.DefaultPostForm("m_rigion", "86")
	for i := range requiredField {
		v := c.PostForm(requiredField[i])
		if v == "" {
			resp.Err(requiredField[i] + " is required")
			return
		}
		m[requiredField[i]] = v
	}
	m["mobile"] = fmt.Sprintf("%s.%s", m["m_rigion"], m["mobile"])

