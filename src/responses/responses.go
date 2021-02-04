package responses

import (
	"async-book-shelf/src/failure"
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if data == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		failure.FailOnError(err, "Unexpected error")
		log.Fatal(err)
	}
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
