package amqp

import (
	"async-book-shelf/src/config"
	"async-book-shelf/src/failure"
)

const (
	emit   = "emit"
	listen = "listen"
)

type AMQPService struct {
	exchangeName string
	exchangeType string
}

func NewRabbitMQService() AMQPService {
	return AMQPService{
		exchangeName: config.RabbitMQExchangeName,
		exchangeType: config.RabbitMQExchangeType,
	}
}

func (service AMQPService) validate(routingKey string) {
	if routingKey == "" {
		failure.Fail("Please provide a routing key")
	}
}

func (service AMQPService) Subscribe(routingKey string, callback func(content []byte)) {
	service.validate(routingKey)
	subscribe(service.exchangeName, service.exchangeType, "", routingKey, callback)
}
