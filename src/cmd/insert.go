package cmd

import (
	"async-book-shelf/src/amqp"
	"async-book-shelf/src/book"
	"async-book-shelf/src/config"
	"async-book-shelf/src/elasticsearch"
	"async-book-shelf/src/failure"
	"encoding/json"
)

func Insert(content string) {
	if content == "" {
		failure.Fail("Please provide a book in JSON format")
	}

	config.Load()

	elasticService := elasticsearch.GetESClient()
	amqpService := amqp.NewRabbitMQService()

	bookService := book.NewBookService(elasticService, amqpService)

	var book book.Book
	err := json.Unmarshal([]byte(content), &book)
	if err != nil {
		failure.FailOnError(err, "Invalid book format")
	}

	err = bookService.Insert(book)
	if err != nil {
		failure.FailOnError(err, "Unable to insert")
	}
}
