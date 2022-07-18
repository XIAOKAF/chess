package service

import (
	"errors"
	"testing"
	"time"
)

func TestInsertCode(t *testing.T) {
	var insertTest = []struct {
		mobile   string
		code     string
		duration time.Duration
	}{
		{"123", "5569", 3},
		{"12345678900", "123456", 5},
		{"12345678900", "123456", 2},
		{"12345678900", "123456", 5},
	}

	for i := 0; i < len(insertTest); i++ {
		test := insertTest[i]
		err := InsertCode(test.mobile, test.code, test.duration)
		if err != nil {
			if err.Error() != errors.New("wrong data").Error() {
				t.Errorf("err:%q", err)
			}
		}
	}
}
