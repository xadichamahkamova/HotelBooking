package handler

import (
	"api-gateway/internal/service"
	"context"
	producer "api-gateway/internal/kafka/producer"
)

type HandlerST struct {
	Service service.ServiceRepositoryClient
	Producer producer.ProducerInit
}

func NewApiHandler(service service.ServiceRepositoryClient, producer producer.ProducerInit) *HandlerST{
	return &HandlerST{
		Service: service,
		Producer: producer,
	}
}
var ctx = context.Background()