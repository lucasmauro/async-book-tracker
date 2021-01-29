package subscriber

import (
	"async-book-shelf/src/config"
	"async-book-shelf/src/failure"
	"log"

	"github.com/streadway/amqp"
)

func Subscribe(exchangeName, exchangeType, queueName, routingKey string) {
	conn, err := amqp.Dial(config.RabbitMQURL)
	failure.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failure.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false, nil)
	failure.FailOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(queueName, false, false, true, false, nil)
	failure.FailOnError(err, "Failed to declare a queue")

	log.Printf("Binding queue '%s' to exchange '%s' with routing key '%s'", q.Name, exchangeName, routingKey)
	err = ch.QueueBind(q.Name, routingKey, exchangeName, false, nil)
	failure.FailOnError(err, "Failed to bind a queue")

	messages, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	failure.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for message := range messages {
			log.Printf(" [R] (Exchange: %s, Routing Key: %s) Received: %s", exchangeName, routingKey, message.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
