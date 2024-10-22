package initialize

import (
	"github.com/devpenguin/go-ecommerce/global"
	"github.com/devpenguin/go-ecommerce/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}