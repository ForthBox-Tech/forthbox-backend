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
func GetAuthInfo(c *gin.Context) {
	resp := NewResp(c)

	token := strings.TrimSpace(c.Query("token"))
	if token == "" {
		resp.Err("token is required")
		return
	}

	cl, err := service.ParseJwtToken(token)
	if err != nil {
		resp.Err(err.Error())
		return
	}

	data := gin.H{"username": cl.Username, "token": token, "expired": cl.ExpiresAt}
	resp.Out(data)
}

// placeholder
func ResetPassword(c *gin.Context) {
	resp := NewResp(c)

	requiredField := []string{"input", "verify_code", "password", "confirm_password"}
	for i := range requiredField {
		f := requiredField[i]
		if strings.TrimSpace(c.PostForm(f)) == "" {
			resp.Err(f + " is required")
			return
		}
	}
	input := strings.TrimSpace(c.PostForm("input"))
	vcode := strings.TrimSpace(c.PostForm("verify_code"))
	password := strings.TrimSpace(c.PostForm("password"))
	cpassword := strings.TrimSpace(c.PostForm("confirm_password"))
	ip := c.ClientIP()
	u, err := service.ResetPassword(input, vcode, password, cpassword, ip)
	if err != nil {
		resp.Err(err.Error())
		return
	}

	data := gin.H{"username": u.UserName}
	resp.Out(data)
}

func SetPassword(c *gin.Context) {
	resp := NewResp(c)
	var password, cpassword string
	password = strings.TrimSpace(c.PostForm("password"))
	if password == "" {
		resp.Err("password is required")
		return
	}
	cpassword = strings.TrimSpace(c.PostForm("confirm_password"))
	if cpassword == "" {
		resp.Err("confirm password is required")
		return
	}

	c.GetHeader("Authorization")
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		resp.Err("Authorization header is empty")
		return
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		resp.Err("Authorization: Bearer token")
		return
	}
	// placeholder
	u, err := service.GetUserFromJwtToken(parts[1])
	if err != nil {
		resp.Err(err.Error())
		return
	}
	// placeholder
	if u.Passwd != "" {
		resp.Err("password not allowed set again")
		return
	}
	if password != cpassword {
		resp.Err("password and confirm password was unmatched")
		return
	}
	// placeholder
	_, err = service.SetUserPassword(u, password)
	if err != nil {
		resp.Err(err.Error())
		return
	}

	data := gin.H{"username": u.UserName}
	resp.Out(data)
}



