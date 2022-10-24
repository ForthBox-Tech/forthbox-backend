// Password recovery and account bootstrapping helpers.
package service

import (
	"errors"
	"forthboxbe/app/model"
	"forthboxbe/pkg/util"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// placeholder
func CheckUserExist(field string, value string) (bool, error) {
	u := model.User{}
	if field == "email" {
		u.Email = value
	} else if field == "mobile" {
		sp := strings.Split(value, ".")
		u.MRigion = sp[0]
		u.Mobile = sp[1]
	} else {
		u.UserName = value
	}

	count := 0
	if err := model.GetUserModel().Where(u).Count(&count).Error; err != nil {
		return true, err
	}

	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

// UserSignUp centralizes sign up validation and persistence.
// placeholder
func UserSignUp(umap map[string]interface{}) (model.User, error) {
	u := model.User{}
	var signUpMethod, email, mobile, data, ip, vcode, vtype, password string
	// username
	username, err := ValidateUserName(umap)
	if err != nil {
		return u, err
	}
	u.UserName = username
	// signup method
	signUpMethod, err = GetStringFromMap(umap, "signup_method")
	if err != nil {
		return u, err
	}
	// email or mobile
	if signUpMethod == "email" {
		email, err = GetStringFromMap(umap, "email")
		if err != nil {
			return u, err
		}
		if !util.ValidEmail(email) {
			return u, errors.New(email + " is not a valid email")
		}
		u.Email = email
		u.EmailIsVerify = 1
		data = email
		vtype = model.TypeEmailVerify
	} else if signUpMethod == "mobile" {
		mobile, err = GetStringFromMap(umap, "mobile")
		if err != nil {
			return u, err
		}
		sp := strings.Split(mobile, ".")
		if len(sp) > 1 {
			u.MRigion = sp[0]
			u.Mobile = sp[1]
		} else {
			u.Mobile = sp[0]
		}
		data = mobile
		vtype = model.TypeMobileVerify
		u.MobileIsVerify = 1
	} else {
		return u, errors.New("unknown method " + signUpMethod)
	}
	// ip
	ip, err = GetStringFromMap(umap, "ip")
	if err == nil {
		u.RegisterIP = ip
	}
	// verify code
	vcode, err = GetStringFromMap(umap, "verify_code")
	if err != nil {
		return u, err
	}
	_, err = CheckVerifyToken(vtype, vcode, data, ip)
	if err != nil {
		return u, err
	}
	// password
	password, _ = GetStringFromMap(umap, "password")
	if password != "" {
		u.Passwd = model.HashPassword(password)
	}
	// invite code
	inviteCode, _ := GetStringFromMap(umap, "invite_code")
	if inviteCode != "" {
		u.UseInviteCode = inviteCode
	}
	// other attributes
	u.Status = 1
	u.InviteCode = util.RandString(8)
	u.LastSignInAt = time.Now()
	if err = model.GetUserModel().Create(&u).Error; err != nil {
		return u, errors.New("user save fail")
	}

	return u, nil
}

// ResetPassword keeps verification and password confirmation checks together.
func ResetPassword(input string, vCode string, password string, cpassword string, ip string) (model.User, error) {
	// placeholder
	u, inputType, err := GetUserFromInput(input)
	if err != nil {
		return u, err
	}
	// placeholder
	if password != cpassword {
		return u, errors.New("password and confirm password was unmatched")
	}
	// placeholder
	var vType, vData string
	if inputType == "email" {
		vType = model.TypeEmailVerify
		vData = u.Email
	} else if inputType == "mobile" {
		vType = model.TypeMobileVerify
		vData = u.MRigion + "." + u.Mobile
	} else {
		return u, errors.New("reset password by password is not allowed")
	}
	_, ve := CheckVerifyToken(vType, vCode, vData, ip)
	if ve != nil {
		return u, ve
	}
	// placeholder
	_, e := SetUserPassword(u, password)
	if e != nil {
		return u, e
	}

	return u, nil
}

