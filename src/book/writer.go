package book

import (
	"async-book-shelf/src/config"
	"context"
	"encoding/json"
	"log"

	"github.com/olivere/elastic/v7"
)

type BookWriter struct {
	elastic *elastic.Client
}

func NewBookWriter(elastic *elastic.Client) BookWriter {
	return BookWriter{
		elastic: elastic,
	}
}

func (writer BookWriter) Publish(message []byte) {
	ctx := context.Background()
	book := string(message)

	_, err := writer.
		elastic.
		Index().
		Index(config.ElasticSearchIndex).
		BodyJson(book).
		Do(ctx)

	if err != nil {
		log.Printf("Unable to publish book: %s\n", err)
	}
}

func (writer BookWriter) Update(message []byte) {
	ctx := context.Background()

	var book Book
	err := json.Unmarshal(message, &book)
	if err != nil {
		log.Printf("Unable to update book: %s\n", err)
		return
	}

	_, err = writer.
		elastic.
		Index().
		Index(config.ElasticSearchIndex).
		Id(book.Id).
		BodyJson(string(message)).
		Do(ctx)

	if err != nil {
		log.Printf("Unable to update book: %s\n", err)
	}
}

func (writer BookWriter) Delete(message []byte) {
	ctx := context.Background()
	bookId := string(message)

	_, err := writer.
		elastic.
		Delete().
		Index(config.ElasticSearchIndex).
		Id(bookId).
		Do(ctx)

	if err != nil {
		log.Printf("Unable to delete book: %s\n", err)
	}
}
