package database

import (
	"buy/app/model"
	"buy/logger"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

var (
	kafkaWriter *kafka.Writer
)

func InitKafka() {
	brokers := viper.GetStringSlice("kafka.brokers")
	if len(brokers) == 0 {
		logger.Panic("缺少kafka broker地址")
	}
	kafkaWriter = &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Topic:        viper.GetString("kafka.topic"),
		Async:        true,
	}
	logger.Info("初始化Kafka成功")
}

func CloseKafka() {
	kafkaWriter.Close()
}

func KafkaWriteOrder(order model.Order) error {
	msg, err := json.Marshal(order)
	if err != nil {
		return err
	}
	return kafkaWriter.WriteMessages(context.Background(), kafka.Message{Value: msg})
}
