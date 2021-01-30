package cmd

import (
	"async-book-shelf/src/config"
	"async-book-shelf/src/controllers"
	"os"
)

const (
	mode       = 1
	routingKey = 2
	content    = 3
)

func getArg(index int) string {
	if len(os.Args) <= index {
		return ""
	}
	return os.Args[index]
}

func Init() {
	config.Load()

	controller := controllers.NewRabbitMQController()
	controller.Mode = getArg(mode)
	controller.RoutingKey = getArg(routingKey)
	controller.Content = getArg(content)

	controller.Validate()
	controller.Run()
}
