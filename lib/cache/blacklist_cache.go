package cache

import (
	"fmt"
	"imserver/lib/redislib"
)

// 黑名单
const (
	backlistPrefix  = "acc:user:backlist" // 数据不重复提交
	UserLoginPrefix = "acc:user:login"    // 用户登陆标记
)

func getBacklistKey() (key string) {
	key = fmt.Sprintf("%s", backlistPrefix)

	return
}

// 设置黑名单用户
func SetUserBacklist(uid int) error {
	key := getBacklistKey()

	redisClient := redislib.GetClient()
	_, err := redisClient.Do("hSet", key, uid, "1").Int()
	return err
}

func CheckBacklist(uid int) (bool, error) {
	key := getBacklistKey()
	redisClient := redislib.GetClient()
	number, err := redisClient.Do("hExists", key, uid).Int()
	if number == 1 {
		return true, err
	}
	return false, err
}

func CheckBackMobilelist(mobile string) (bool, error) {
	key := getBacklistKey()
	redisClient := redislib.GetClient()
	number, err := redisClient.Do("hExists", key, mobile).Int()
	if number == 1 {
		return true, err
	}
	return false, err
}
