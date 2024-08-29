package consumer

import (
	"fmt"
	config "websocket-service/internal/pkg/load"
	"websocket-service/logger"

	"github.com/twmb/franz-go/pkg/kgo"
)

type IConsumeInit interface {
	ConsumeMessage() error
	Close() error
}

type ConsumeInit struct {
	Client *kgo.Client
	Topic  string
}

func NewConsumeInit(cfg *config.Config) (*ConsumeInit, error) {

	address := fmt.Sprintf("%s:%s", cfg.Kafka.Host, cfg.Kafka.Port)
	client, err := kgo.NewClient(
		kgo.SeedBrokers(address),
		kgo.ConsumeTopics(cfg.Kafka.Topic),
	)
	if err != nil {
		logger.Error("Failed to create Kafka client: ", err)
		return nil, err
	}

	return &ConsumeInit{Client: client, Topic: cfg.Kafka.Topic}, nil
}

func (p *ConsumeInit) Close() error {
	p.Client.Close()
	return nil
}
