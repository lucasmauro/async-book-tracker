package book

import (
	"async-book-shelf/src/config"
	"context"
	"encoding/json"

	"github.com/olivere/elastic/v7"
)

type BookReader struct {
	elastic *elastic.Client
}

func NewBookReader(elastic *elastic.Client) BookReader {
	return BookReader{
		elastic: elastic,
	}
}

func (reader BookReader) Get(key string, value interface{}) ([]Book, error) {
	ctx := context.Background()

	var books []Book

	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery(key, value))

	searchService := reader.elastic.Search().Index(config.ElasticSearchIndex).SearchSource(searchSource)

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
