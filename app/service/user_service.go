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
