package dao

import (
	"chess/proto"
)

func SelectMobile(user *proto.RegisterRequest) (string, error) {
	var mobile string
	result := DB.Table("user").Select("mobile").Where("mobile = ?", user.Mobile)
	result.Scan(&mobile)
	return mobile, result.Error
}

func InsertUser(user *proto.RegisterRequest) error {
	result := DB.Select("mobile", "password").Create(&user)
	return result.Error
}

func SelectPwd(user *proto.LoginRequest) (string, error) {
	u := &proto.LoginRequest{}
	result := DB.Select("password").Where("mobile = ?", user.Mobile).Scan(&u)
	return u.Password, result.Error
}
