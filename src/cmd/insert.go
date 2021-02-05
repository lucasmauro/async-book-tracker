package cmd

import (
	"async-book-shelf/src/amqp"
	"async-book-shelf/src/book"
	"async-book-shelf/src/config"
	"async-book-shelf/src/failure"
	"encoding/json"
	"fmt"
)

func Insert(content string) {
	if content == "" {
		failure.Fail("Please provide a book in JSON format")
	}

	config.Load()

	amqpService := amqp.NewRabbitMQService()

	bookWriter := book.NewBookWriter(amqpService)

	var book book.Book
	fmt.Println(content)
	err := json.Unmarshal([]byte(content), &book)
	if err != nil {
		failure.FailOnError(err, "Invalid book format")
	}

	err = bookWriter.Insert(book)
	if err != nil {
		failure.FailOnError(err, "Unable to insert")
	}
}
