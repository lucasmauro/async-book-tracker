package controllers

import (
	"async-book-shelf/src/book"
	"async-book-shelf/src/elasticsearch"
	"async-book-shelf/src/responses"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	elastic := elasticsearch.GetESClient()

	reader := book.NewBookReader(elastic)

	query := r.URL.Query()
	key := query.Get("key")
	value := query.Get("value")

	books, err := reader.Get(key, value)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, books)
}
