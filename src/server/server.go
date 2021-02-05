package server

import (
	"async-book-shelf/src/config"
	"async-book-shelf/src/failure"
	"async-book-shelf/src/router"
	"fmt"
	"net/http"
)

func Serve() {
	r := router.Generate()

	fmt.Printf("Listening on port %s\n", config.ServerPort)

	port := ":" + config.ServerPort
	err := http.ListenAndServe(port, r)
	failure.FailOnError(err, "Unable to serve")
}
