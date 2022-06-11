package server_service

import (
	"chess/server/dao"
	"chess/server/model"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

func ParseSmsConfig() (model.Message, error) {
	var s model.Message
	file, err := os.Open("config/Sms.json")
	if err != nil {
		return s, err
	}
	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		return s, err
	}
	err = json.Unmarshal(fileByte, &s)
	if err != nil {
		return s, err
	}
	return s, nil
}

func InsertCode(mobile string, code string, duration time.Duration) error {
	err := dao.Set(mobile, code, duration)
	if err != nil {
		return err
	}
	return nil
}

func GetCode(mobile string) (string, error) {
	code, err := dao.Get(mobile + "code")
	if err != nil {
		return code, err
	}
	return code, nil
}
