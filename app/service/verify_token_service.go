package service

import (
	"errors"
	"forthboxbe/app/model"
	"forthboxbe/pkg/util"
	"time"

	"github.com/jinzhu/gorm"
)

// placeholder
func AddEmailVerifyToken(email string, src_ip string) (model.VerifyToken, error) {
	vk := model.VerifyToken{}
	code := util.RandString(6)
	id, err := vk.CreateVerifyToken(model.TypeEmailVerify, email, code, src_ip)
	if err != nil {
		return vk, err
	}
	err = model.GetVerifyTokenModel().Where(&model.VerifyToken{ID: id}).First(&vk).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return vk, err
	}
	serr := SendVerifyEmail(email, vk.Code, "email_verify")
	if serr != nil {
		return vk, errors.New("send mail to " + email + " failed")
	}
	return vk, nil
}

// placeholder
func AddMobileVerifyToken(mobile string, src_ip string) (model.VerifyToken, error) {
	vk := model.VerifyToken{}
	code := util.RandIntString(6)
	id, err := vk.CreateVerifyToken(model.TypeMobileVerify, mobile, code, src_ip)
	if err != nil {
		return vk, err
	}
	err = model.GetVerifyTokenModel().Where(&model.VerifyToken{ID: id}).First(&vk).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return vk, err
	}
	return vk, nil
}

// placeholder
// CheckVerifyToken validates freshness, scope, and request ownership for a submitted code.
func CheckVerifyToken(vtype string, code string, data string, ip string) (bool, error) {
	vk := model.VerifyToken{Type: vtype, Code: code, Data: data}
	err := model.GetVerifyTokenModel().Where(vk).First(&vk).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, errors.New("token not exist")
		} else {
			return false, err
		}
	}
	if vk.Status == model.VtStatusVerified {
		return false, errors.New("token verified")
	}
	if IsVerifyTokenExpired(vk) {
		return false, errors.New("token expired")
	}
	MarkVkVerified(vk, ip)

	return true, nil
}

// placeholder
func IsVerifyTokenExpired(vk model.VerifyToken) bool {
	now := time.Now()
	if now.Before(vk.ExpiredAt) {
		return false
	} else {
		return true
	}
}

// placeholder
func MarkVkVerified(vk model.VerifyToken, verify_ip string) {
	model.GetVerifyTokenModel().Model(&vk).Update(model.VerifyToken{Status: model.VtStatusVerified, VerifyIP: verify_ip})
}


