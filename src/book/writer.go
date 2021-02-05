package book

import (
	"async-book-shelf/src/amqp"
	"async-book-shelf/src/config"
	"encoding/json"
)

type BookWriter struct {
	amqp amqp.AMQPService
}

func NewBookWriter(amqp amqp.AMQPService) BookWriter {
	return BookWriter{
		amqp: amqp,
	}
}

func (writer BookWriter) Insert(book Book) error {
	content, err := json.Marshal(book)
	if err != nil {
		return err
	}

	writer.amqp.Publish(config.RabbitMQInsertionRoutingKey, string(content))
	return nil
}
