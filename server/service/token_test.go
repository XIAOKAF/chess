package service

import (
	"errors"
	"testing"
	"time"
)

func TestCreateToken(t *testing.T) {
	var createTest = []struct {
		mobile      string
		duration    time.Duration
		tokenString string
		err         error
	}{
		{"12345678900", 2, "12345678900*", nil},
		{"123", 2, "", errors.New("wrong mobile")},
	}
	for _, tt := range createTest {
		err, _ := CreateToken(tt.mobile, tt.duration)
		if err != nil {
			if err.Error() != tt.err.Error() {
				t.Fatalf("err:%q", err)
			}
		}
	}
}
