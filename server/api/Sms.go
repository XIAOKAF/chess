package api

import (
	"chess/server/server_service"
	"chess/service"
	"chess/tool"
	"context"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tencentsms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"math/rand"
	"time"
)

func (u *userService) Send(ctx context.Context, request *service.SendRequest) (*service.Response, error) {
	var resp *service.Response
	if request.Mobile == "" {
		response := tool.Failure(resp, 400, "必要字段不能为空")
		return response, nil
	}
	//生成随机验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	sms, err := server_service.ParseSmsConfig()
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		fmt.Println("1", err)
		return response, err
	}
	credential := common.NewCredential(sms.SecretId, sms.SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, err := tencentsms.NewClient(credential, "ap-guangzhou", cpf)
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		fmt.Println("2", err)
		return response, err
	}
	req := tencentsms.NewSendSmsRequest()
	req.SmsSdkAppId = common.StringPtr(sms.AppId)
	req.SignName = common.StringPtr(sms.Sign)
	req.SenderId = common.StringPtr("")
	req.ExtendCode = common.StringPtr("")
	req.TemplateParamSet = common.StringPtrs([]string{code, "5"})
	req.TemplateId = common.StringPtr(sms.TemplateId)
	req.PhoneNumberSet = common.StringPtrs([]string{"+86" + request.Mobile})
	//发送短信
	_, err = client.SendSms(req)
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		fmt.Println("3", err)
		return response, err
	}
	//存储短信
	err = server_service.InsertCode(request.Mobile+"code", code, 5)
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		fmt.Println("4", err)
		return response, err
	}
	response := tool.Success(resp, 200, "短信发送成功")
	return response, nil
}
