package middleware

import (
	"imserver/common"
	"imserver/lib/cache"
	"imserver/lib/jwt"

	"github.com/gin-gonic/gin"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(200, common.Response(common.ParameterIllegal, "", nil))
			c.Abort()
			return
		}

		// 验证token
		uid, _, isValid := jwt.ParseToken(token)
		if !isValid || uid <= 0 {
			c.JSON(200, common.Response(common.ErrorTokenExpire, "", nil))
			c.Abort()
			return
		}

		// 验证redis中是否存在这个token,如果不存在，则跳转到登陆页面

		// cache.SetUserBacklist(uid)

		// 验证uid是否在黑名单中
		flag, err := cache.CheckBacklist(uid)
		if err != nil {
			c.JSON(200, common.Response(common.ServerError, "", nil))
			c.Abort()
			return
		}
		if flag { // 在黑名单中
			c.JSON(200, common.Response(common.UserLoginNotAllow, "", nil))
			c.Abort()
			return
		}

		// c.Header("token", newToken) // 刷新后的token
		c.Set("uid", uid)
	}
}
