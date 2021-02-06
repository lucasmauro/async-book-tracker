package amqp

import (
	"async-book-shelf/src/config"
	"async-book-shelf/src/failure"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func getConnection() (*amqp.Connection, error) {
	conn, err := amqp.Dial(config.RabbitMQURL)

	if err != nil {
		for start := time.Now(); err != nil && time.Since(start) < time.Second*config.RabbitMQTimeout; {
			log.Println("Failed to connect to RabbitMQ. Retrying...")
			time.Sleep(time.Second * 3)
			conn, err = amqp.Dial(config.RabbitMQURL)
		}
		if err != nil {
			return nil, err
		}
	}

	return conn, nil
}

func subscribe(exchangeName, exchangeType, queueName, routingKey string, callback func(content []byte)) {
	conn, err := getConnection()
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
			callback(message.Body)
		}
	}()

	log.Printf(" [*] (Exchange: %s, Routing Key: %s) Waiting for messages. To exit press CTRL+C", exchangeName, routingKey)
	<-forever
}
