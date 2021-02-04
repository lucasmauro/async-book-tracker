package config

import (
	"async-book-shelf/src/failure"
	"fmt"
	"os"
	"strings"
)

const (
	server_port = "SERVER_PORT"
)

var ServerPort = ""

type server struct {
	port string
}

func (config *server) load() {
	config.port = os.Getenv(server_port)
}

func (config *server) validate() {
	var invalidVariables []string

	if config.port == "" {
		invalidVariables = append(invalidVariables, server_port)
	}

	if len(invalidVariables) > 0 {
		variables := strings.Join(invalidVariables, ", ")
		message := fmt.Sprintf("Empty environment variable(s): [%s]", variables)
		failure.Fail(message)
	}
}

func loadServer() {
	var config server
	config.load()
	config.validate()

	ServerPort = config.port
}
