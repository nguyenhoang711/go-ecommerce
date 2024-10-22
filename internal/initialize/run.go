package initialize

import "github.com/gin-gonic/gin"

func Run() *gin.Engine{

	LoadConfig()
	InitLogger()

	r := InitRouter()
	// r.GET("../../docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
