package controller

import (
	"chess/proto"
	"context"
	"errors"
	tencentsms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"testing"
)

func TestUserService_Send(t *testing.T) {
	var ctx context.Context
	var sendTest = []struct {
		ctx     context.Context
		mobile  string
		code    int32
		message string
		err     error
	}{
		{ctx, "", 400, "必要字段不能为空", nil},
		{ctx, "456", 400, "", errors.New(tencentsms.FAILEDOPERATION_PHONENUMBERPARSEFAIL)}, //电话号码格式错误
		{ctx, "18725917586", 200, "短信发送成功", nil},
		{ctx, "18725917586", 500, "", errors.New(tencentsms.FAILEDOPERATION_PHONENUMBERINBLACKLIST)}, //黑名单电话
		{ctx, "18725917586", 500, "", errors.New(tencentsms.FAILEDOPERATION_CONTAINSENSITIVEWORD)},   //敏感信息
	}
	for _, tt := range sendTest {
		req := &proto.SendRequest{
			Mobile: tt.mobile,
		}
		resp, err := UserService.Send(tt.ctx, req)
		if err != nil {
			if err.Error() != tt.err.Error() {
				t.Errorf("err:%q", err)
			}
			if resp.Code != tt.code {
				t.Errorf("want code:%d,get:%d", tt.code, resp.Code)
			}
			if resp.Info != tt.message {
				t.Errorf("want info:%s,get:%s", tt.message, resp.Info)
			}
		}
	}
}
