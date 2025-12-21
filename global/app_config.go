package global

import (
	"database/sql"

	"github.com/devpenguin/go-ecommerce/pkg/logger"
	"github.com/devpenguin/go-ecommerce/pkg/settings"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

var (
	Config        settings.Config
	Logger        *logger.LoggerZap
	Mdb           *sql.DB
	Rdb           *redis.Client
	KafkaProducer *kafka.Writer
)
