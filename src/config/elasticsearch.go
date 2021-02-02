package config

import (
	"async-book-shelf/src/failure"
	"fmt"
	"os"
	"strings"
)

const (
	elasticsearch_address  = "ELASTICSEARCH_ADDRESS"
	elasticsearch_port     = "ELASTICSEARCH_PORT"
	elasticsearch_user     = "ELASTICSEARCH_USER"
	elasticsearch_password = "ELASTICSEARCH_PASSWORD"
)

var ElasticSearchURL = ""

type elasticsearch struct {
	address  string
	port     string
	user     string
	password string
}

func (config *elasticsearch) load() {
	config.address = os.Getenv(elasticsearch_address)
	config.port = os.Getenv(elasticsearch_port)
	config.user = os.Getenv(elasticsearch_user)
	config.password = os.Getenv(elasticsearch_password)
}

func (config *elasticsearch) validate() {
	var invalidVariables []string

	if config.address == "" {
		invalidVariables = append(invalidVariables, elasticsearch_address)
	}

	if config.port == "" {
		invalidVariables = append(invalidVariables, elasticsearch_port)
	}

	if config.user == "" {
		invalidVariables = append(invalidVariables, elasticsearch_user)
	}

	if config.password == "" {
		invalidVariables = append(invalidVariables, elasticsearch_password)
	}

	if len(invalidVariables) > 0 {
		variables := strings.Join(invalidVariables, ", ")
		message := fmt.Sprintf("Empty environment variable(s): [%s]", variables)
		failure.Fail(message)
	}
}

func loadElasticSearch() {
	var config elasticsearch
	config.load()
	config.validate()

	ElasticSearchURL = fmt.Sprintf(
		"http://%s:%s@%s:%s",
		config.user,
		config.password,
		config.address,
		config.port,
	)
}
