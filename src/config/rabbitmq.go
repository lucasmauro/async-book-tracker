package config

import (
	"async-book-shelf/src/failure"
	"fmt"
	"os"
	"strings"
)

const (
	address      = "RABBITMQ_ADDRESS"
	port         = "RABBITMQ_PORT"
	user         = "RABBITMQ_USER"
	password     = "RABBITMQ_PASSWORD"
	exchangeName = "RABBITMQ_EXCHANGE_NAME"
	exchangeType = "RABBITMQ_EXCHANGE_TYPE"
)

var RabbitMQURL = ""
var RabbitMQExchangeName = ""
var RabbitMQExchangeType = ""

type rabbitMQ struct {
	address      string
	port         string
	user         string
	password     string
	exchangeName string
	exchangeType string
}

func load(config *rabbitMQ) {
	config.address = os.Getenv(address)
	config.port = os.Getenv(port)
	config.user = os.Getenv(user)
	config.password = os.Getenv(password)
	config.exchangeName = os.Getenv(exchangeName)
	config.exchangeType = os.Getenv(exchangeType)
}

func validate(config *rabbitMQ) {
	var invalidVariables []string

	if config.address == "" {
		invalidVariables = append(invalidVariables, address)
	}

	if config.port == "" {
		invalidVariables = append(invalidVariables, port)
	}

	if config.user == "" {
		invalidVariables = append(invalidVariables, user)
	}

	if config.password == "" {
		invalidVariables = append(invalidVariables, password)
	}

	if config.exchangeName == "" {
		invalidVariables = append(invalidVariables, exchangeName)
	}

	if config.exchangeType == "" {
		invalidVariables = append(invalidVariables, exchangeType)
	}

	if len(invalidVariables) > 0 {
		variables := strings.Join(invalidVariables, ", ")
		message := fmt.Sprintf("Empty environment variable(s): [%s]", variables)
		failure.Fail(message)
	}
}

func loadRabbitMQ() {
	var config rabbitMQ
	load(&config)
	validate(&config)

	RabbitMQURL = fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		config.user,
		config.password,
		config.address,
		config.port,
	)
	RabbitMQExchangeName = config.exchangeName
	RabbitMQExchangeType = config.exchangeType
}
