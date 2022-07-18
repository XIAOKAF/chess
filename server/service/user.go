package service

import (
	"chess/proto"
	"chess/server/dao"
)

func SelectUser(user *proto.RegisterRequest) (error, bool) {
	_, err := dao.SelectMobile(user)
	return err, true
}

func InsertUser(user *proto.RegisterRequest) error {
	err := dao.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}

func SelectPwd(user *proto.LoginRequest) (string, error) {
	return dao.SelectPwd(user)
}
