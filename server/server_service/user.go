package server_service

import (
	"chess/server/dao"
	"chess/service"
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

func InsertUser(user *service.RegisterRequest) error {
	err := dao.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}
