package initialize

import (
	"net/http"

	"github.com/devpenguin/go-ecommerce/global"
	"github.com/devpenguin/go-ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default() //tao 1 instance cua gin (middleware, version, etc ...)
	//use the middleware
	// r.Use(AA(), BB(), CC)
	r.GET("/health-check", Check)
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", Pong)          // /v1/2024/ping
	}

	return r
}

// middleware
// func AA() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("Before --> AA")
// 		c.Next()
// 		fmt.Println("After --> AA")
// 	}
// }

// func BB() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("Before --> BB")
// 		c.Next()
// 		fmt.Println("After --> BB")
// 	}
// }
// func CC(c *gin.Context) {
// 	fmt.Println("Before --> CC")
// 	c.Next()
// 	fmt.Println("After --> CC")
// }

func Pong(c *gin.Context) {
	// name := c.Param("name")
	uid := c.Query("uid") //query params
	//default query
	name := c.DefaultQuery("name", "anonystick")
	c.JSON(http.StatusOK, gin.H{ //map string in java
		// "message": "pong...ping" + name,
		"message": "Hello, " + name,
		"uid":     uid,
	})
}

func Check(c *gin.Context) {
	response.SuccessResponse(c, http.StatusOK, "Health check", gin.H{
		"Server Port": global.Config.Server.Port,
		"Mode": global.Config.Server.Mode,
	})
}