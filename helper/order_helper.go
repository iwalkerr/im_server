/**
* Created by GoLand.
* User: link1st
* Date: 2019-08-01
* Time: 18:13
 */

package helper

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

func GetOrderIdTime() (orderId string) {

	currentTime := time.Now().Nanosecond()
	orderId = fmt.Sprintf("%d", currentTime)

	return
}

// 随机生成图片
func RandGenImage() string {
	curDir, err := os.Getwd()
	if err == nil {
		curDir += "/"
	}

	files, _ := ioutil.ReadDir(curDir + "assets/default")
	if len(files) == 0 {
		return ""
	}
	rand.Seed(time.Now().Unix())
	number := rand.Intn(len(files))
	numberString := strconv.Itoa(number)

	httpUrl := viper.GetString("app.httpUrl")
	return fmt.Sprintf(httpUrl+"/default/%s.jpeg", numberString)
}
