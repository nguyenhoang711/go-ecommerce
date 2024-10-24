package initialize

import (
	"github.com/devpenguin/go-ecommerce/global"
	repos "github.com/devpenguin/go-ecommerce/internal/models"
	"github.com/devpenguin/go-ecommerce/internal/service"
	"github.com/devpenguin/go-ecommerce/internal/service/impl"
)

func InitServiceInterface() {
	queries := repos.New(global.Mdb)

	service.InitUserAuth(impl.NewUserAuthImpl(queries))

	global.Logger.Info("Init all service interface success!")
}