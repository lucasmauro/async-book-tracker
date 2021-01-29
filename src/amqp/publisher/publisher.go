package publisher

import (
	"async-book-shelf/src/config"
	"async-book-shelf/src/failure"
	"log"

	"github.com/streadway/amqp"
)

func Publish(exchangeName, exchangeType, routingKey, content string) {
	conn, err := amqp.Dial(config.RabbitMQURL)
	failure.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failure.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false, nil)
	failure.FailOnError(err, "Failed to declare an exchange")

	publishing := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(content),
	}
	err = ch.Publish(exchangeName, routingKey, false, false, publishing)
	failure.FailOnError(err, "Failed to publish a message")

	log.Printf(" [S] (Exchange: %s, Routing Key: %s) Sent: %s", exchangeName, routingKey, content)
}
