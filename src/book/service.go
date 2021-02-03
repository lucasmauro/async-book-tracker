package book

import (
	"async-book-shelf/src/amqp"
	"async-book-shelf/src/config"
	"context"
	"encoding/json"

	"github.com/olivere/elastic/v7"
)

type BookService struct {
	elastic *elastic.Client
	amqp    amqp.AMQPService
}

func NewBookService(elastic *elastic.Client, amqp amqp.AMQPService) BookService {
	return BookService{
		elastic: elastic,
		amqp:    amqp,
	}
}

func (service BookService) Get(key string, value interface{}) ([]Book, error) {
	ctx := context.Background()

	var books []Book

	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery(key, value))

	searchService := service.elastic.Search().Index(config.ElasticSearchIndex).SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		return books, err
	}

	for _, hit := range searchResult.Hits.Hits {
		var book Book
		err := json.Unmarshal(hit.Source, &book)
		if err != nil {
			return books, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (service BookService) Insert(book Book) error {
	content, err := json.Marshal(book)
	if err != nil {
		return err
	}

	service.amqp.Publish(config.RabbitMQInsertionRoutingKey, string(content))
	return nil
}
