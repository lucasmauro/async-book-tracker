package cmd

import (
	"async-book-shelf/src/amqp"
	"async-book-shelf/src/book"
	"async-book-shelf/src/config"
	"async-book-shelf/src/elasticsearch"
)

func Init() {
	config.Load()

	elasticService := elasticsearch.GetESClient()
	amqpService := amqp.NewRabbitMQService()

	bookService := book.NewBookService(elasticService, amqpService)
	bookService.Get("name", "The Lord of the Rings: The Return of the King")
}
