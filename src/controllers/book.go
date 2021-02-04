package controllers

import (
	"async-book-shelf/src/amqp"
	"async-book-shelf/src/book"
	"async-book-shelf/src/elasticsearch"
	"async-book-shelf/src/responses"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	elastic := elasticsearch.GetESClient()
	amqp := amqp.NewRabbitMQService()
	service := book.NewBookService(elastic, amqp)

	query := r.URL.Query()
	key := query.Get("key")
	value := query.Get("value")

	books, err := service.Get(key, value)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, books)
}
