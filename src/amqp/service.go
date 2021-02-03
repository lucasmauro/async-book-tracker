package amqp

import (
	"async-book-shelf/src/amqp/publisher"
	"async-book-shelf/src/amqp/subscriber"
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

func (service AMQPService) Publish(routingKey, content string) {
	service.validate(routingKey)
	publisher.Publish(service.exchangeName, service.exchangeType, routingKey, content)
}

func (service AMQPService) Subscribe(routingKey string) {
	service.validate(routingKey)
	subscriber.Subscribe(service.exchangeName, service.exchangeType, "", routingKey)
}
