/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-25
* Time: 12:20
 */

package routers

import (
	swag "imserver/controllers/swagger"
	"imserver/controllers/systems"
	"imserver/controllers/user"
	"imserver/middleware"
	"imserver/swagger"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	// router.LoadHTMLGlob("views/**/*")

	router.GET("/tool/swagger", swag.Swagger)
	swagger.Init(router) // 初始化文档

	// 获取验证码
	router.POST("/authCode", user.SendSMS)

	// 中间件，验证是否拉入黑名单
	router.Use()

	// 用户登陆
	router.POST("/login", user.Login)
	// 用户注册
	router.POST("/register", user.Register)
	// 忘记密码
	router.POST("/forgetPwd", user.ForgetPwd)

	// 中间件,拦截登陆
	router.Use(middleware.CheckToken())

	// 用户组
	userRouter := router.Group("/user")
	userRouter.POST("/integral", user.UserIntegral)       // 获取积分
	userRouter.POST("/headImg", user.UpdateHeadImg)       // 更新头像
	userRouter.POST("/userInfo", user.GetUserInfo)        // 获取用户信息
	userRouter.POST("/updateUser", user.UpdateUserInfo)   // 更新用户信息
	userRouter.POST("/friend", user.AddFriend)            // 添加朋友
	userRouter.POST("/editPwd", user.EditUserPwd)         // 修改用户密码
	userRouter.POST("/bandingAlipay", user.BandingAlipay) // 绑定alipay
	userRouter.POST("/getAlipay", user.GetAlipayInfo)     // 获取alipay信息
	userRouter.POST("/addAddr", user.AddShopAddr)         // 添加购物地址
	userRouter.POST("/editAddr", user.EditShopAddr)       // 修改购物地址
	userRouter.POST("/addrList", user.ShopAddrList)       // 购物地址列表
	userRouter.GET("/list", user.List)
	userRouter.GET("/online", user.Online)
	userRouter.POST("/sendMessage", user.SendMessage)
	userRouter.POST("/sendMessageAll", user.SendMessageAll)

	// 系统
	systemRouter := router.Group("/system")
	systemRouter.GET("/state", systems.Status)

	// home
	// homeRouter := router.Group("/home")
	// homeRouter.GET("/index", home.Index)

	// 上传图片、语音、视频

}
