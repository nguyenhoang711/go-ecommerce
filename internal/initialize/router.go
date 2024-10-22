package initialize

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default() //tao 1 instance cua gin (middleware, version, etc ...)
	//use the middleware
	r.Use(AA(), BB(), CC)
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", Pong)          // /v1/2024/ping
		// v1.DELETE("/ping", controller.NewPongController().Pong)
		// v1.PATCH("/ping", controller.NewPongController().Pong)
		// v1.HEAD("/ping", controller.NewPongController().Pong)
		// v1.OPTIONS("/ping", controller.NewPongController().Pong)
	}
	return r

}

// middleware
func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next()
		fmt.Println("After --> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("After --> BB")
	}
}
func CC(c *gin.Context) {
	fmt.Println("Before --> CC")
	c.Next()
	fmt.Println("After --> CC")
}

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
