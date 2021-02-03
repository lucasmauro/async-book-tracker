package config

import (
	"async-book-shelf/src/failure"
	"fmt"
	"os"
	"strings"
)

const (
	rabbitmq_address             = "RABBITMQ_ADDRESS"
	rabbitmq_port                = "RABBITMQ_PORT"
	rabbitmq_user                = "RABBITMQ_USER"
	rabbitmq_password            = "RABBITMQ_PASSWORD"
	rabbitmq_exchangeName        = "RABBITMQ_EXCHANGE_NAME"
	rabbitmq_exchangeType        = "RABBITMQ_EXCHANGE_TYPE"
	rabbitmq_insertionRoutingKey = "RABBITMQ_INSERTION_ROUTING_KEY"
)

var RabbitMQURL = ""
var RabbitMQExchangeName = ""
var RabbitMQExchangeType = ""
var RabbitMQInsertionRoutingKey = ""

type rabbitMQ struct {
	address             string
	port                string
	user                string
	password            string
	exchangeName        string
	exchangeType        string
	insertionRoutingKey string
}

func (config *rabbitMQ) load() {
	config.address = os.Getenv(rabbitmq_address)
	config.port = os.Getenv(rabbitmq_port)
	config.user = os.Getenv(rabbitmq_user)
	config.password = os.Getenv(rabbitmq_password)
	config.exchangeName = os.Getenv(rabbitmq_exchangeName)
	config.exchangeType = os.Getenv(rabbitmq_exchangeType)
	config.insertionRoutingKey = os.Getenv(rabbitmq_insertionRoutingKey)
}

func (config *rabbitMQ) validate() {
	var invalidVariables []string

	if config.address == "" {
		invalidVariables = append(invalidVariables, rabbitmq_address)
	}

	if config.port == "" {
		invalidVariables = append(invalidVariables, rabbitmq_port)
	}

	if config.user == "" {
		invalidVariables = append(invalidVariables, rabbitmq_user)
	}

	if config.password == "" {
		invalidVariables = append(invalidVariables, rabbitmq_password)
	}

	if config.exchangeName == "" {
		invalidVariables = append(invalidVariables, rabbitmq_exchangeName)
	}

	if config.exchangeType == "" {
		invalidVariables = append(invalidVariables, rabbitmq_exchangeType)
	}

	if config.insertionRoutingKey == "" {
		invalidVariables = append(invalidVariables, rabbitmq_insertionRoutingKey)
	}

	if len(invalidVariables) > 0 {
		variables := strings.Join(invalidVariables, ", ")
		message := fmt.Sprintf("Empty environment variable(s): [%s]", variables)
		failure.Fail(message)
	}
}

func loadRabbitMQ() {
	var config rabbitMQ
	config.load()
	config.validate()

	RabbitMQURL = fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		config.user,
		config.password,
		config.address,
		config.port,
	)
	RabbitMQExchangeName = config.exchangeName
	RabbitMQExchangeType = config.exchangeType
	RabbitMQInsertionRoutingKey = config.insertionRoutingKey
}
