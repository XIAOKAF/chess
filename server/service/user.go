package service

import (
	"chess/proto"
	"chess/server/dao"
	"database/sql"
	"fmt"
)

func SelectUser(mobile string) (error, bool) {
	err := dao.SelectMobile(mobile)
	fmt.Println(err)
	if err == sql.ErrNoRows {
		return nil, false
	}
	return err, true
}

func InsertUser(user *proto.RegisterRequest) error {
	err := dao.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}

func SelectPwd(mobile string) (string, error) {
	return dao.SelectPwd(mobile)
}
