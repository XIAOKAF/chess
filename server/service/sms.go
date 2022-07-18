package service

import (
	"chess/server/dao"
	"chess/server/model"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

func ParseSmsConfig() (model.Message, error) {
	var s model.Message
	file, err := os.Open("config/" + "sms.json")
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
	if len(mobile) != 11 || len(code) != 6 || duration != 2*time.Minute {
		return errors.New("wrong data")
	}
	err := dao.Set(mobile, code, duration)
	if err != nil {
		return err
	}
	return nil
}

func GetCode(mobile string) (string, error) {
	if len(mobile) != 11 {
		return "", errors.New("wrong mobile")
	}
	code, err := dao.Get(mobile + "code")
	return code, err
}
