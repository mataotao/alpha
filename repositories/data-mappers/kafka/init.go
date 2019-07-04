package kafka

import (
	"github.com/Shopify/sarama"
)

func (Config Config) Init() (sarama.Client, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	kafkaClient, err := sarama.NewClient(Config.Addrs, config)
	return kafkaClient, err
}
