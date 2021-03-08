package helper

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/go-kratos/kratos/pkg/log"
)

// todo: 替换为公司账号
const (
	SignName          = "结草"
	TemplateCode      = "SMS_203020216"
	REGION_ID         = "cn-shenzhen"
	ACCESS_KEY_ID     = "LTAI4GEvSQJEGC2wPhsmsbA2"
	ACCESS_KEY_SECRET = "wJCYVMG7S4PfNqSWL45J61CxRH0Bh3"
)

// 发送验证码
func SendSms(phone string, code string) (err error) {
	client, err := sdk.NewClientWithAccessKey(REGION_ID, ACCESS_KEY_ID, ACCESS_KEY_SECRET)
	if err != nil {
		log.Error("ali ecs client failed, err:%s", err.Error())
		return
	}

	request := requests.NewCommonRequest()                           // 构造一个公共请求
	request.Method = "POST"                                          // 设置请求方式
	request.Product = "Ecs"                                          // 指定产品
	request.Scheme = "https"                                         // https | http
	request.Domain = "dysmsapi.aliyuncs.com"                         // 指定域名则不会寻址，如认证方式为 Bearer Token 的服务则需要指定
	request.Version = "2017-05-25"                                   // 指定产品版本
	request.ApiName = "SendSms"                                      // 指定接口名
	request.QueryParams["RegionId"] = "cn-hangzhou"                  // 地区
	request.QueryParams["PhoneNumbers"] = phone                      //手机号
	request.QueryParams["SignName"] = SignName                       //阿里云验证过的项目名 自己设置
	request.QueryParams["TemplateCode"] = TemplateCode               //阿里云的短信模板号 自己设置
	request.QueryParams["TemplateParam"] = "{\"code\":" + code + "}" //短信模板中的验证码内容 自己生成

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		log.Error("ali ecs client failed, err:%s", err.Error())
		return
	}
	log.Info(response.String())
	var message Message //阿里云返回的json信息对应的类
	//记得判断错误信息
	_ = json.Unmarshal(response.GetHttpContentBytes(), &message)
	if message.Message != "OK" {
		// 错误处理
		return
	}
	return nil
}

//json数据解析
type Message struct {
	Message   string
	RequestId string
	BizId     string
	Code      string
}

//生成指定位数的数字
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

//生成指定位数的字符
func GenerateSubId(width int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, width)
	for i := range b {
		b[i] = letterRunes[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(letterRunes))]
	}
	return string(b)
}

// 生成前缀
func RandUUID() string {
	var letterRunes = []rune("abcdefghjklmnopqrstuvwxyzABCDEFGHJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 3)
	for i := range b {
		b[i] = letterRunes[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(letterRunes))]
	}
	return string(b) + randSub()
}

// 生成后缀
func randSub() string {
	var letterRunes = []rune("0123456789")
	b := make([]rune, 5)
	for i := range b {
		b[i] = letterRunes[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(letterRunes))]
	}

	return string(b)
}
