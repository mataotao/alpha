package kafka

import (
	"alpha/config"
	"github.com/Shopify/sarama"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

func Send(client sarama.Client, topic, key string, notify interface{}) error {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	valueBytes, err := json.Marshal(notify)
	if err != nil {
		return err
	}
	message := sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(valueBytes),
	}

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return err
	}

	defer func() {
		if err := producer.Close(); err != nil {
		}
	}()

	partition, offset, err := producer.SendMessage(&message)
	if err != nil {
		return err
	}
	config.Logger.Info("kafka send",
		zap.String("topic", topic),
		zap.Any("message", message),
	)
	config.Logger.Info("send message to kafka succeed!",
		zap.String("topic", topic),
		zap.Any("key", key),
		zap.Any("partition", partition),
		zap.Any("offset", offset),
	)
	return nil
}
