package main

import (
	"github.com/devpenguin/go-ecommerce/internal/initialize"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/devpenguin/go-ecommerce/cmd/swag/docs"
)

// @title           Swagger GO Ecommerce API
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  github.com/devpenguin/ecommerce

// @contact.name   Tips Go
// @contact.url    github.com/devpenguin/ecommerce
// @contact.email  hoangnguyendang777@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8002
// @BasePath  /v1/2024
// @schema http
func main() {
	r := initialize.Run()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8002")
}