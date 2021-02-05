package book

import (
	"async-book-shelf/src/config"
	"context"
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