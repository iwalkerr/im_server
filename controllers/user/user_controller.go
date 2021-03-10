/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-25
* Time: 12:11
 */

package user

import (
	"fmt"
	"imserver/common"
	"imserver/controllers"
	"imserver/helper"
	"imserver/lib/cache"
	"imserver/lib/jwt"
	"imserver/lib/redislib"
	"imserver/models"
	"imserver/repository"
	"imserver/servers/websocket"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm

const (
	userSmsPrefix    = "user:sms:" // 用户在线状态
	userSmsCacheTime = 5 * 60      // 5分钟过期
)

func GetCode(key string) string {
	return "123456"

	client := redislib.GetClient()
	code := client.Get(key).Val()
	return code
}

// 添加朋友
func AddFriend(c *gin.Context) {

}

// 获取所有群和单人聊天信息
func GetAllTeam(c *gin.Context) {
	uid := c.GetInt("uid")
	list := repository.GetAllTeamByUid(uid)
	c.JSON(200, common.Response(common.OK, "", list))
}

// @Tags 用户管理
// @Title 删除收货地址
// @Summary 删除收货地址
// @Description 删除收货地址
// @Accept multipart/form-data
// @Produce json
// @Param token header string true "用户token" default(eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm)
// @Param shop_addr_id formData string true "用户收货地址ID" default(1)
// @Success 200 {object} common.JsonResult
// @Router /user/deleteAddr [post]
func DeleteAddr(c *gin.Context) {
	shopAddrId := c.PostForm("shop_addr_id")
	if shopAddrId == "" || shopAddrId == "0" {
		c.JSON(200, common.Response(common.ParameterIllegal, "", nil))
		return
	}
	uid := c.GetInt("uid")
	if err := repository.DeleteShopAddr(uid, shopAddrId); err != nil {
		c.JSON(200, common.Response(common.ServerError, "", nil))
		return
	}
	c.JSON(200, common.Response(common.OK, "", nil))
}

// 添加收货地址
// @Tags 用户管理
// @Title 添加收货地址
// @Summary 添加收货地址
// @Description 添加收货地址
// @Accept multipart/form-data
// @Produce json
// @Param token header string true "用户token" default(eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm)
// @Param name formData string true "姓名" default(张三)
// @Param phone formData string true "联系电话" default(13881887710)
// @Param addr formData string true "地址" default(四川/成都/温江区)
// @Param detail formData string true "详细地址" default(南薰大道塞纳湖畔7栋503)
// @Param is_default formData int true "是否默认，1默认" default(0)
// @Success 200 {object} common.JsonResult
// @Router /user/addAddr [post]
func AddShopAddr(c *gin.Context) {
	var req models.ShopAddr
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, common.Response(common.ParameterIllegal, "", nil))
		return
	}

	req.Uid = c.GetInt("uid")

	// 如果地址是默认地址，则取消数据库中其他为默认的地址
	if req.IsDefault == 1 {
		if err := repository.CancelDefaultAddr(req.Uid); err != nil {
			c.JSON(200, common.Response(common.ServerError, "", nil))
			return
		}
	}
	// 添加地址
	if err := repository.AddUserShopAddr(&req); err != nil {
		c.JSON(200, common.Response(common.ServerError, "", nil))
		return
	}

	c.JSON(200, common.Response(common.OK, "", nil))
}

// 修改收货地址
// @Tags 用户管理
// @Title 修改收货地址
// @Summary 修改收货地址
// @Description 修改收货地址
// @Accept multipart/form-data
// @Produce json
// @Param token header string true "用户token" default(eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm)
// @Param shop_addr_id formData string true "用户收货地址ID" default(1)
// @Param name formData string true "姓名" default(张三)
// @Param phone formData string true "联系电话" default(13881887710)
// @Param addr formData string true "地址" default(四川/成都/温江区)
// @Param detail formData string true "详细地址" default(南薰大道塞纳湖畔7栋503)
// @Param is_default formData int true "是否默认，1默认" default(0)
// @Success 200 {object} common.JsonResult
// @Router /user/editAddr [post]
func EditShopAddr(c *gin.Context) {
	var req models.EditShopAddr
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, common.Response(common.ParameterIllegal, "", nil))
		return
	}

	req.Uid = c.GetInt("uid")
	if req.IsDefault == 1 {
		if err := repository.CancelDefaultAddr(req.Uid); err != nil {
			c.JSON(200, common.Response(common.ServerError, "", nil))
			return
		}
	}

	// 修改用户地址
	if err := repository.EditUserAddr(&req); err != nil {
		c.JSON(200, common.Response(common.ServerError, "", nil))
		return
	}

	c.JSON(200, common.Response(common.OK, "", nil))
}

// 获取收货地址列表
// @Tags 用户管理
// @Title 获取收货地址列表
// @Summary 获取收货地址列表
// @Description 获取收货地址列表
// @Accept multipart/form-data
// @Produce json
// @Param token header string true "用户token" default(eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm)
// @Success 200 {object} common.JsonResult
// @Router /user/addrList [post]
func ShopAddrList(c *gin.Context) {
	uid := c.GetInt("uid")
	list := repository.ShopAddrList(uid)
	c.JSON(200, common.Response(common.OK, "", list))
}

// 绑定alipay
// @Tags 用户管理
// @Title 绑定alipay
// @Summary 绑定alipay
// @Description 绑定alipay
// @Accept multipart/form-data
// @Produce json
// @Param token header string true "用户token" default(eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm)
// @Param name formData string true "用户名" default(张三)
// @Param account formData string true "账号" default(320553500@qq.com)
// @Success 200 {object} common.JsonResult
// @Router /user/bandingAlipay [post]
func BandingAlipay(c *gin.Context) {
	realname := c.PostForm("name")
	alipayAccount := c.PostForm("account")
	if realname == "" || alipayAccount == "" {
		c.JSON(200, common.Response(common.ParameterIllegal, "", nil))
		return
	}
	// todo 阿里验证信息

	uid := c.GetInt("uid")
	if err := repository.UpdateAlipayInfo(uid, realname, alipayAccount); err != nil {
		c.JSON(200, common.Response(common.ServerError, "", nil))
		return
	}

	c.JSON(200, common.Response(common.OK, "", nil))
}

// @Tags 用户管理
// @Title 获取alipay信息
// @Summary 获取alipay
// @Description 获取alipay
// @Accept multipart/form-data
// @Produce json
// @Param token header string true "用户token" default(eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm)
// @Success 200 {object} common.JsonResult
// @Router /user/getAlipay [post]
func GetAlipayInfo(c *gin.Context) {
	uid := c.GetInt("uid")
	realname, alipayAccount := repository.GetAlipayInfo(uid)
	c.JSON(200, common.Response(common.OK, "", gin.H{"realname": realname, "alipay_account": alipayAccount}))
}

// 修改用户密码
// @Tags 用户管理
// @Title 修改用户密码
// @Summary 修改用户密码
// @Description 修改用户密码
// @Accept multipart/form-data
// @Produce json
// @Param token header string true "用户token" default(eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm)
// @Param old_pwd formData string true "就密码" default(123456)
// @Param new_pwd formData string true "新密码" default(3221)
// @Success 200 {object} common.JsonResult
// @Router /user/editPwd [post]
func EditUserPwd(c *gin.Context) {
	var req models.PwdInfo
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, common.Response(common.ParameterIllegal, "", nil))
		return
	}
	// 验证老密码的争取性
	uid := c.GetInt("uid")
	resp, err := repository.GetPwdById(uid)
	if err != nil {
		c.JSON(200, common.Response(common.ServerError, "", nil))
		return
	}

	pwd := resp.Mobile + req.OldPwd + resp.Signature
	pwd = helper.MD5(pwd)
	if resp.Password != pwd {
		c.JSON(200, common.Response(common.OldPwdError, "", nil))
		return
	}

	// 更新新密码
	pwd = resp.Mobile + req.NewPwd + resp.Signature
	pwd = helper.MD5(pwd)
	if err := repository.UpdatePwd(resp.Mobile, pwd); err != nil {
		c.JSON(200, common.Response(common.ServerError, "", nil))
		return
	}

	c.JSON(200, common.Response(common.OK, "", nil))
}

// 获取用户信息
// @Tags 用户管理
// @Title 获取用户信息
// @Summary 获取用户信息
// @Description 获取用户信息
// @Accept multipart/form-data
// @Produce json
// @Param token header string true "用户token" default(eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm)
// @Success 200 {object} common.JsonResult
// @Router /user/userInfo [post]
func GetUserInfo(c *gin.Context) {
	uid := c.GetInt("uid")

	userInfo := repository.GetUserInfo(uid)
	c.JSON(200, common.Response(common.OK, "", userInfo))
}

// 更新用户信息
// @Tags 用户管理
// @Title 更新用户信息
// @Summary 更新用户信息
// @Description 更新用户信息
// @Accept multipart/form-data
// @Produce json
// @Param token header string true "用户token" default(eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm)
// @Param name formData string true "昵称" default(张三)
// @Param sex formData string false "性别" default(1男/2女)
// @Param email formData string false "邮箱" default(320553500@qq.com)
// @Param user_addr formData string false "用户地址" default(四川省/成都市/成华区)
// @Success 200 {object} common.JsonResult
// @Router /user/updateUser [post]
func UpdateUserInfo(c *gin.Context) {
	var req models.UserInfo
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, common.Response(common.ParameterIllegal, "", nil))
		return
	}

	req.Uid = c.GetInt("uid")
	if err := repository.UpdateUser(&req); err != nil {
		c.JSON(200, common.Response(common.ServerError, "", nil))
		return
	}

	c.JSON(200, common.Response(common.OK, "", nil))
}

// @Tags 用户管理
// @Title 更新用户头像
// @Summary 更新用户头像
// @Description 更新用户头像
// @Accept multipart/form-data
// @Produce json
// @Param token header string true "用户token" default(eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm)
// @Param img formData file true "用户头像"
// @Success 200 {object} common.JsonResult
// @Router /user/headImg [post]
func UpdateHeadImg(c *gin.Context) {
	uid := c.GetInt("uid")
	header, err := c.FormFile("img")
	if err != nil {
		c.JSON(200, common.Response(common.ServerError, "", nil))
		return
	}

	curDir, _ := os.Getwd()
	curdate := time.Now().UnixNano()
	dts := curDir + "/assets/head/"

	ext := path.Ext(header.Filename)
	filename := strconv.Itoa(uid) + strconv.FormatInt(curdate, 10) + ext
	dts = dts + filename

	if err := c.SaveUploadedFile(header, dts); err != nil {
		c.JSON(200, common.Response(common.ServerError, "", nil))
		return
	}

	httpUrl := viper.GetString("app.httpUrl")
	imgpath := httpUrl + "/head/" + filename

	// 存入数据库
	if err := repository.UploadImage(uid, imgpath); err != nil {
		c.JSON(200, common.Response(common.ServerError, "", nil))
		return
	}

	c.JSON(200, common.Response(common.OK, "", imgpath))
}

// 用户积分
// @Tags 用户管理
// @Title 用户积分
// @Summary 用户积分
// @Description 用户积分
// @Accept x-www-form-urlencoded
// @Produce json
// @Param token header string true "用户token" default(eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYxNzA5MzgxNCwiaXNzIjoid2ViX2dpbl90ZW1wbGF0ZSIsIlJlZnJlc2hUaW1lIjoxNjE1Nzk3ODE0fQ.LmTyXCemtbMu7rGcxQtefEdU04RdVbp9hk2_6ulW8retAULYhVa5zvRBl8IXoFEm)
// @Success 200 {object} common.JsonResult
// @Router /user/integral [post]
func UserIntegral(c *gin.Context) {
	uid := c.GetInt("uid")

	integral := repository.GetUserIntegral(uid)
	c.JSON(200, common.Response(common.OK, "", integral))
}

// @Tags 手机验证码
// @Title 发送手机验证码
// @Summary 发送手机验证码
// @Description 发送手机验证码
// @Accept x-www-form-urlencoded
// @Produce json
// @Param mobile formData string true "手机号码" default(13881887710)
// @Param type formData string true "发送类型:1登陆,2注册,3忘记密码" default(1)
// @Success 200 {object} common.JsonResult
// @Router /authCode [post]
func SendSMS(c *gin.Context) {
	mobile := c.PostForm("mobile")
	types := c.PostForm("type") // 1登陆，2注册，3忘记密码

	// 验证手机号码的合法性
	if !helper.Validate(mobile) {
		c.JSON(200, common.Response(common.PhoneError, "", nil))
		return
	}

	// 验证是否是60秒内发送的
	_, ok := cache.Instance().Get(mobile)
	if ok {
		c.JSON(200, common.Response(common.SendTimeError, "", nil))
		return
	}

	// 通过第三方发送短信
	code := helper.GenValidateCode(6)
	if err := helper.SendSms(mobile, code); err != nil {
		c.JSON(200, common.Response(common.SendTimeError, "", nil))
		return
	}

	// 验证码间隔一分钟发一次
	cache.Instance().Set(mobile, code, 1*time.Minute)

	// 发送验证码后，存入redis,hash key手机号码，value 6位数验证码，过期时间5分钟

	client := redislib.GetClient()
	// 设置过期时间
	var key string
	if types == "1" {
		key = userSmsPrefix + "login_" + mobile
	} else if types == "2" {
		key = userSmsPrefix + "register_" + mobile
	} else {
		key = userSmsPrefix + "forget_" + mobile
	}
	_, err := client.Do("setEx", key, userSmsCacheTime, code).Result()
	if err != nil {
		c.JSON(200, common.Response(common.ServerError, "", nil))
		return
	}

	c.JSON(200, common.Response(common.OK, "", nil))
}

// 用户登陆
// 短信验证 1分钟一次，5分钟内有效
// @Tags 用户登陆
// @Title 用户登陆
// @Summary 用户登陆
// @Description 用户登陆
// @Accept x-www-form-urlencoded
// @Produce json
// @Param mobile formData string true "手机号码" default(13881887710)
// @Param password formData string true "密码" default(123456)
// @Param sms_code formData string true "短信验证码" default(432343)
// @Param timestamp formData string true "时间戳" default(1614479454)
// @Param sign formData string true "签名" default(5fae1a0e4156e1ad4e742f4b9b7c1416)
// @Success 200 {object} common.JsonResult
// @Router /login [post]
func Login(c *gin.Context) {
	var req models.LoginReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, common.Response(common.ParameterIllegal, "", nil))
		return
	}

	// redis中查询是否有客户端登陆

	// 1. 验证手机号码的合法性
	if !helper.Validate(req.Mobile) {
		c.JSON(200, common.Response(common.PhoneError, "", nil))
		return
	}

	// 2. 是否在黑名单中
	if ok, _ := cache.CheckBackMobilelist(req.Mobile); ok {
		c.JSON(200, common.Response(common.Backlist, "", nil))
		return
	}

	// 3. 验证短信码是否正确，60秒内重新发送
	// 根据手机号码到redis中查询验证码
	key := userSmsPrefix + "login_" + req.Mobile
	code := GetCode(key)
	if req.SmsCode != code {
		c.JSON(200, common.Response(common.AuthCodeError, "", nil))
		return
	}

	// 4. 超过60秒，则过期
	ts := time.Now().Unix()
	if (req.Timestamp + 60) < ts {
		c.JSON(200, common.Response(common.ErrorTokenExpire, "", nil))
		return
	}

	// 验证手机号是否一样，防止重放
	dataMap := map[string]interface{}{
		"mobile":    req.Mobile,
		"password":  req.Password,
		"sms_code":  req.SmsCode,
		"timestamp": req.Timestamp,
		"salt":      "XXXXXXXXXXXX", // 私钥，不用于网络传输
	}

	// 5. 通过传输的数据，验证sign
	sign := helper.GetSign(dataMap)
	if sign != req.Sign {
		c.JSON(200, common.Response(common.ErrorTokenExpire, "", nil))
		return
	}

	// 6. 验证用户名密码
	// 获取用户salt
	user, err := repository.Login(req.Mobile)
	if err != nil {
		c.JSON(200, common.Response(common.ServerError, "", nil))
		return
	}

	pwd := req.Mobile + req.Password + user.Signature
	pwd = helper.MD5(pwd)

	if user.Password != pwd {
		c.JSON(200, common.Response(common.ErrorTokenExpire, "", nil))
		return
	}

	// 7. 生成token
	token, err := jwt.GenerateToken(user.UserId)
	if err != nil {
		c.JSON(200, common.Response(common.ServerError, "", nil))
		return
	}

	// 返回token和用户信息
	c.JSON(200, common.Response(common.OK, "", gin.H{
		"token":        token,
		"head_picture": user.HeadPicture,
		"uuid":         user.Username,
		"name":         user.Name,
		"qrcode":       user.QrCode,
	}))
}

// 用户注册
// @Tags 用户登陆
// @Title 用户注册
// @Summary 用户注册
// @Description 用户注册
// @Accept x-www-form-urlencoded
// @Produce json
// @Param mobile formData string true "手机号码" default(13881887710)
// @Param password formData string true "密码" default(123456)
// @Param sms_code formData string true "短信验证码" default(432343)
// @Param name formData string true "用户名" default(张三))
// @Param sex formData string true "性别:1男,2女" default(1))
// @Param timestamp formData string true "时间戳" default(1614479454)
// @Param sign formData string true "签名" default(5fae1a0e4156e1ad4e742f4b9b7c1416)
// @Success 200 {object} common.JsonResult
// @Router /register [post]
func Register(c *gin.Context) {
	var req models.RegisterReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, common.Response(common.ParameterIllegal, "", nil))
		return
	}

	// 1. 验证手机号码的合法性
	if !helper.Validate(req.Mobile) {
		c.JSON(200, common.Response(common.PhoneError, "", nil))
		return
	}

	// 2. 验证短信码是否正确，60秒内重新发送
	// 根据手机号码到redis中查询验证码
	key := userSmsPrefix + "register_" + req.Mobile
	code := GetCode(key)
	if req.SmsCode != code {
		c.JSON(200, common.Response(common.AuthCodeError, "", nil))
		return
	}

	// 3. 超过60秒，则过期
	ts := time.Now().Unix()
	if (req.Timestamp + 60) < ts {
		c.JSON(200, common.Response(common.ErrorTokenExpire, "", nil))
		return
	}

	// 验证手机号是否一样，防止重放
	dataMap := map[string]interface{}{
		"mobile":    req.Mobile,
		"password":  req.Password,
		"sms_code":  req.SmsCode,
		"name":      req.Name,
		"sex":       req.Sex,
		"timestamp": req.Timestamp,
		"salt":      "XXXXXXXXXXXX", // 私钥，不用于网络传输
	}

	// 4. 通过传输的数据，验证sign
	sign := helper.GetSign(dataMap)
	if sign != req.Sign {
		c.JSON(200, common.Response(common.ErrorTokenExpire, "", nil))
		return
	}

	// 5. 存入数据库
	// 手机号是否已经注册
	if ok := repository.CheckPhone(req.Mobile); ok {
		c.JSON(200, common.Response(common.ExistPhone, "", nil))
		return
	}
	// 生成密码
	salt := helper.GenerateSubId(6)
	pwd := req.Mobile + req.Password + salt
	pwd = helper.MD5(pwd)

	// 生成UUID
	req.Username = GetUUID()
	req.HeadPicture = helper.RandGenImage()
	req.Salt = salt
	req.Password = pwd

	id, err := repository.Register(&req)
	if err != nil {
		c.JSON(200, common.Response(common.ServerError, "", nil))
		return
	}

	// 生成个人二维码
	param := fmt.Sprintf(`{"url":"/user/friend","param":%d}`, id)
	dest := fmt.Sprintf("assets/qrcode/%s.png", req.Username)
	_ = helper.CreateQRCodeByBoombuler(param, qr.M, 200, dest)

	// 二维码存入数据库
	httpUrl := viper.GetString("app.httpUrl")
	qrpath := httpUrl + "/qrcode/" + req.Username + ".png"
	// 更新二维码路径
	if err := repository.UpdateQrcode(id, qrpath); err != nil {
		c.JSON(200, common.Response(common.ServerError, "", nil))
		return
	}

	c.JSON(200, common.Response(common.OK, "", nil))
}

func GetUUID() string {
	uuid := helper.RandUUID()
	// 数据库去查
	if ok := repository.ExistByUUID(uuid); !ok {
		return uuid
	}
	return GetUUID()
}

// 忘记密码
// 用户登陆
// 短信验证 1分钟一次，5分钟内有效
// @Tags 用户登陆
// @Title 忘记密码
// @Summary 忘记密码
// @Description 忘记密码
// @Accept x-www-form-urlencoded
// @Produce json
// @Param mobile formData string true "手机号码" default(13881887710)
// @Param password formData string true "密码" default(123456)
// @Param sms_code formData string true "短信验证码" default(432343)
// @Param timestamp formData string true "时间戳" default(1614479454)
// @Param sign formData string true "签名" default(5fae1a0e4156e1ad4e742f4b9b7c1416)
// @Success 200 {object} common.JsonResult
// @Router /forgetPwd [post]
func ForgetPwd(c *gin.Context) {
	var req models.LoginReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, common.Response(common.ParameterIllegal, "", nil))
		return
	}

	// 1. 验证手机号码的合法性
	if !helper.Validate(req.Mobile) {
		c.JSON(200, common.Response(common.PhoneError, "", nil))
		return
	}

	// 2. 是否在黑名单中
	if ok, _ := cache.CheckBackMobilelist(req.Mobile); ok {
		c.JSON(200, common.Response(common.Backlist, "", nil))
		return
	}

	// 3. 验证短信码是否正确，60秒内重新发送
	// 根据手机号码到redis中查询验证码
	key := userSmsPrefix + "forget_" + req.Mobile
	code := GetCode(key)
	if req.SmsCode != code {
		c.JSON(200, common.Response(common.AuthCodeError, "", nil))
		return
	}

	// 4. 超过60秒，则过期
	ts := time.Now().Unix()
	if (req.Timestamp + 60) < ts {
		c.JSON(200, common.Response(common.ErrorTokenExpire, "", nil))
		return
	}

	// 验证手机号是否一样，防止重放
	dataMap := map[string]interface{}{
		"mobile":    req.Mobile,
		"password":  req.Password,
		"sms_code":  req.SmsCode,
		"timestamp": req.Timestamp,
		"salt":      "XXXXXXXXXXXX", // 私钥，不用于网络传输
	}

	// 5. 通过传输的数据，验证sign
	sign := helper.GetSign(dataMap)
	if sign != req.Sign {
		c.JSON(200, common.Response(common.ErrorTokenExpire, "", nil))
		return
	}

	// 6. 验证用户名密码
	// 获取用户salt
	user, err := repository.Login(req.Mobile)
	if err != nil {
		c.JSON(200, common.Response(common.UserNoRegister, "", nil))
		return
	}

	pwd := req.Mobile + req.Password + user.Signature
	pwd = helper.MD5(pwd)

	// 更新数据库
	if err := repository.UpdatePwd(req.Mobile, pwd); err != nil {
		c.JSON(200, common.Response(common.ServerError, "", nil))
		return
	}

	c.JSON(200, common.Response(common.OK, "", nil))
}

type SendMessageData struct {
	TeamId  int    `json:"teamId" binding:"required"`
	MsgType int    `json:"msgType" binding:"required"`
	Message string `json:"message" binding:"required"`
}

// 给全员发送消息
func SendMessageAll(c *gin.Context) {
	// 获取参数
	var req SendMessageData
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, common.Response(common.ParameterIllegal, "", nil))
		return
	}
	userId := c.GetInt("uid")
	// fmt.Println("http_request 给全体用户发送消息", appIdStr, userId, msgId, message)

	data := make(map[string]interface{})
	// if cache.SeqDuplicates(msgId) {
	// 	fmt.Println("给用户发送消息 重复提交:", msgId)
	// 	controllers.Response(c, common.OK, "", data)
	// 	return
	// }

	sendResults, err := websocket.SendUserMessageAll(req.TeamId, userId, req.MsgType, models.MessageCmdMsg, req.Message)
	if err != nil {
		data["sendResultsErr"] = err.Error()
	}

	data["sendResults"] = sendResults

	controllers.Response(c, common.OK, "", data)
}
