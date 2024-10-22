package global

import (
	"github.com/devpenguin/go-ecommerce/pkg/logger"
	"github.com/devpenguin/go-ecommerce/pkg/settings"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
)