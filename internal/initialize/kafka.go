package initialize

import (
	"log"

	"github.com/devpenguin/go-ecommerce/global"
	"github.com/segmentio/kafka-go"
)

// init kafka producer
var KafkaProducer *kafka.Writer

func InitKafka() {
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "otp-auth-topic",
		Balancer: &kafka.LeastBytes{},
	}
}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatal("Failed to close kafka producer:  %v", err)
	}
}
