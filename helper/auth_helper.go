package helper

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"sort"
	"strconv"
)

const (
	regular = "^(13[0-9]|14[0-9]|15[0-9]|18[0-9])\\d{8}$"
)

func Validate(mobileNum string) bool {
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// 生成32位MD5
func MD5(text string) string {
	h := md5.New()
	_, _ = h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

// 加密原则
// 1.请求来源的合法性、
// 2.请求参数被攥改
// 3.请求的唯一性(不可复制)(加时间戳60秒后过期)
func GetSign(jsonMap map[string]interface{}) string {
	// 获取key值
	var i int
	addressKey := make([]string, len(jsonMap))
	for k := range jsonMap {
		addressKey[i] = k
		i++
	}

	// 根据key的ASCII字符串排序
	sort.Strings(addressKey)

	// 字符串拼接
	var str string
	for _, k := range addressKey {
		switch t := jsonMap[k].(type) {
		case string:
			str += k + t
		case int:
			str += k + strconv.Itoa(t)
		case int64:
			str += k + strconv.FormatInt(t, 10)
		default:
			println("不支持的类型")
		}
	}
	// md5加密
	return MD5(str)
}
