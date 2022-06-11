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

func SelectPwd(mobile string) (string, error) {
	var pwd string
	rows := MDB.QueryRow("SELECT password from user WhERE mobile = ?", mobile)
	if rows.Err() != nil {
		return "", rows.Err()
	}
	err := rows.Scan(&pwd)
	if err != nil {
		return "", err
	}
	return pwd, nil
}
