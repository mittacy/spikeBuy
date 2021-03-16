package database

import (
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
	"goods/logger"
)

func NewKafkaReader(groupId string) *kafka.Reader {
	brokers := viper.GetStringSlice("kafka.brokers")
	if len(brokers) == 0 {
		logger.Panic("缺少kafka broker")
	}

	topic := viper.GetString("kafka.topic")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupId,
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
}