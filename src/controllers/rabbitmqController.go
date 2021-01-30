package controllers

import (
	"async-book-shelf/src/amqp/publisher"
	"async-book-shelf/src/amqp/subscriber"
	"async-book-shelf/src/config"
	"async-book-shelf/src/failure"
	"fmt"
)

const (
	emit = "emit"
	listen = "listen"
)

type RabbitMQController struct {
	Mode         string
	ExchangeName string
	ExchangeType string
	RoutingKey   string
	Content      string
}

func NewRabbitMQController() RabbitMQController {
	return RabbitMQController{
		ExchangeName: config.RabbitMQExchangeName,
		ExchangeType: config.RabbitMQExchangeType,
	}
}

func (controller RabbitMQController) Validate() {
	if (controller.Mode == "") || (controller.Mode != emit && controller.Mode != listen) {
		failure.Fail(fmt.Sprintf("Please provide a valid running mode [%s|%s]", emit, listen))
	}

	if controller.RoutingKey == "" {
		failure.Fail("Please provide a routing key")
	}

	if controller.Mode == emit && controller.Content == "" {
		failure.Fail("Please provide a content")
	}
}

func (controller RabbitMQController) Run() {
	if controller.Mode == emit {
		publisher.Publish(controller.ExchangeName, controller.ExchangeType, controller.RoutingKey, controller.Content)
		return
	}

	if controller.Mode == listen {
		subscriber.Subscribe(controller.ExchangeName, controller.ExchangeType, "", controller.RoutingKey)
		return
	}
}
