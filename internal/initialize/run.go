package initialize

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine{
	env := os.Getenv("ENV")
	if env == "" {
        log.Fatal("ENV variable is not set")
		LoadConfig("local")
    } else {
		LoadConfig(env)
	}
	InitLogger()

	InitMysqlc()
	InitRedis()
	InitServiceInterface()

	r := InitRouter()

	return r
}
