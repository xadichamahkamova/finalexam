package producer

import (
	"incexp-service/logger"
	"incexp-service/internal/pkg/load"
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

type IProducerInit interface {
	ProduceMessage(key string, message []byte) error
	Close() error
}

type ProducerInit struct {
	Client *kgo.Client
	Topic  string
}


func NewProducerInit(cfg *load.Config) (*ProducerInit, error) {

	address := fmt.Sprintf("%s:%s", cfg.Kafka.Host, cfg.Kafka.Port)
	logger.Info("Initializing Kafka producer with address:", address)
	client, err := kgo.NewClient(
		kgo.SeedBrokers(address),
		kgo.AllowAutoTopicCreation(),
	)
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}
	logger.Info("Kafka client initialized successfully")
	return &ProducerInit{
		Client: client, 
		Topic: cfg.Kafka.Topic,
	}, nil
}

func (p *ProducerInit) ProduceMessage(key string, message []byte) error {

	logger.Info("Producing message to topic:", p.Topic, " with key:", key)
	
	record := &kgo.Record{
		Topic: p.Topic,
		Key:   []byte(key),
		Value: message,
	}
	
	err := p.Client.ProduceSync(context.Background(), record).FirstErr()
	if err != nil {
		logger.Error(err)
		return err
	}
	logger.Info("Message produced successfully to topic:", p.Topic)
	return nil
}

func (p *ProducerInit) Close() error {

	logger.Info("Closing Kafka client")
	p.Client.Close()
	return nil
}
