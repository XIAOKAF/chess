package dao

import "chess/service"

func SelectMobile(mobile string) error {
	rows := MDB.QueryRow("SELECT password from user WHERE mobile = ?", mobile)
	return rows.Err()
}

func InsertUser(user *service.RegisterRequest) error {
	result := DB.Select("mobile", "password").Create(&user)
	return result.Error
}
