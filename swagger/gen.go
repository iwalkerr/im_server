package swagger

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// API文档
func Init(router *gin.Engine) {
	host := viper.GetString("app.httpUrl")

	SwaggerInfo.Title = "Golang API 测试"
	SwaggerInfo.Version = "2.0"
	SwaggerInfo.Description = fmt.Sprintf(`生成文档请在调试模式下进行 <a href=\"http://%s/tool/swagger?a=r\">重新生成文档</a>`, host)
	SwaggerInfo.Host = host
	SwaggerInfo.BasePath = "/"
	SwaggerInfo.Schemes = []string{"http", "https"}

	g1 := router.Group("/swagger")
	g1.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://"+host+"/swagger/doc.json")))
}
