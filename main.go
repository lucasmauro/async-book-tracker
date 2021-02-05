package main

import (
	"async-book-shelf/src/amqp"
	"async-book-shelf/src/book"
	"async-book-shelf/src/config"
	"async-book-shelf/src/elasticsearch"
	"async-book-shelf/src/server"
	"os"
)

func getArg(index int) string {
	if len(os.Args) <= index {
		return ""
	}
	return os.Args[index]
}

func main() {
	config.Load()

	amqp := amqp.NewRabbitMQService()
	elastic := elasticsearch.GetESClient()
	writer := book.NewBookWriter(elastic)

	go amqp.Subscribe(config.RabbitMQInsertionRoutingKey, writer.Publish)

	server.Serve()
}
