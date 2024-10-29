package initialize

import (
	"net/http"

	"github.com/devpenguin/go-ecommerce/internal/controller/auth"
	"github.com/devpenguin/go-ecommerce/internal/controller/syst"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default() //tao 1 instance cua gin (middleware, version, etc ...)
	//use the middleware
	// r.Use(AA(), BB(), CC)
	v1 := r.Group("/syst")
	{
		v1.GET("/ping", Pong)          // /v1/2024/ping
		v1.POST("/system-info", syst.SystemAPI.Ping)
	}

	authAPI := r.Group("/auth")
	{
		authAPI.POST("/register", auth.UserAuth.Register)
		authAPI.POST("/verify_account", auth.UserAuth.VerifyOTP)
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