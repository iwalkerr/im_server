/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-25
* Time: 12:11
 */

package common

const (
	OK                 = 200  // Success
	NotLoggedIn        = 1000 // 未登录
	ParameterIllegal   = 1001 // 参数不合法
	UnauthorizedUserId = 1002 // 非法的用户Id
	Unauthorized       = 1003 // 未授权
	ServerError        = 1004 // 系统错误
	NotData            = 1005 // 没有数据
	ModelAddError      = 1006 // 添加错误
	ModelDeleteError   = 1007 // 删除错误
	ModelStoreError    = 1008 // 存储错误
	OperationFailure   = 1009 // 操作失败
	RoutingNotExist    = 1010 // 路由不存在
	ErrorTokenExpire   = 1011 // token过期
	UserLoginNotAllow  = 1012 // 用户暂时不允许登录
	AuthCodeError      = 1013 // 验证码错误
	PhoneError         = 1014 // 手机号格式错误
	SendTimeError      = 1015 // 60秒发一次
	Backlist           = 1016 // 用户在黑名单中
	ExistPhone         = 1017 // 手机号码已经注册
	UserNoRegister     = 1018 // 该用户没有注册
	OldPwdError        = 1019 // 原密码错误
)

// 根据错误码 获取错误信息
func GetErrorMessage(code uint32, message string) string {
	var codeMessage string
	codeMap := map[uint32]string{
		OK:                 "Success",
		NotLoggedIn:        "未登录",
		ParameterIllegal:   "参数不合法",
		UnauthorizedUserId: "非法的用户Id",
		Unauthorized:       "未授权",
		NotData:            "没有数据",
		ServerError:        "系统错误",
		ModelAddError:      "添加错误",
		ModelDeleteError:   "删除错误",
		ModelStoreError:    "存储错误",
		OperationFailure:   "操作失败",
		RoutingNotExist:    "路由不存在",
		ErrorTokenExpire:   "token过期",
		UserLoginNotAllow:  "用户暂时不允许登录",
		AuthCodeError:      "验证码错误",
		PhoneError:         "手机号格式错误",
		SendTimeError:      "60秒发一次",
		Backlist:           "用户在黑名单中",
		ExistPhone:         "手机号已经注册",
		UserNoRegister:     "用户没有注册",
		OldPwdError:        "原密码错误",
	}

	if message == "" {
		if value, ok := codeMap[code]; ok {
			// 存在
			codeMessage = value
		} else {
			codeMessage = "未定义错误类型!"
		}
	} else {
		codeMessage = message
	}

	return codeMessage
}
