// Authentication services coordinate password checks and token based user lookups.
package service

import (
	"errors"
	"forthboxbe/app/model"
	"time"
)

// placeholder
func UserSignIn(input string, password string, ip string) (model.User, error) {
	u, _, err := GetUserFromInput(input)
	if err != nil {
		return u, err
	}

	if !u.CheckPassword(password) {
		return u, errors.New("password not matched")
	}
	model.GetUserModel().Model(&u).Update(model.User{LastSignInAt: time.Now(), LastSignInIP: ip})

	return u, nil
}

// placeholder
func GetUserFromJwtToken(jwt string) (model.User, error) {
	u := model.User{}
	cl, err := ParseJwtToken(jwt)
	if err != nil {
		return u, err
	}
	u.UserName = cl.Username
	err = model.GetUserModel().Where(u).First(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}



