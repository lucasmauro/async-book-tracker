package config

import (
	"async-book-shelf/src/failure"

	"github.com/joho/godotenv"
)

func Load() {
	if err := godotenv.Load(); err != nil {
		failure.FailOnError(err, "Unable to load godotenv")
	}

	loadServer()
	loadRabbitMQ()
	loadElasticSearch()
}
