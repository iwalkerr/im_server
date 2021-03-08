package swagger

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

// http://47.108.146.201:8080/swagger/index.html
// swagger文档
func Swagger(c *gin.Context) {
	if a := c.Query("a"); a == "r" {
		curDir, err := os.Getwd()
		if err == nil {
			curDir += "/"
		}
		genPath := curDir + "swagger"

		if _, err := ExecCommand("swag init -o " + genPath); err != nil { //重新生成文档
			log.Printf("运行命令行 ===> 错误提示: %s", err.Error())
		} else {
			log.Println("文档重新编译成功")
		}
	}
	c.Redirect(http.StatusFound, "/swagger/index.html")
}

// 运行cmd
func ExecCommand(arg string) (string, error) {
	cmd := exec.Command("bash", "-c", arg)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out //剧透，坑在这里
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		// log.Printf("运行命令行 %s ===> 错误提示: %s", err.Error(), stderr.String())
		return "", err
	}
	return out.String(), nil
}
