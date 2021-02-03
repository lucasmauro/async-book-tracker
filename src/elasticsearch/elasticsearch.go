package elasticsearch

import (
	"async-book-shelf/src/config"
	"async-book-shelf/src/failure"

	"github.com/olivere/elastic/v7"
)

func GetESClient() *elastic.Client {
	client, err := elastic.NewClient(elastic.SetURL(config.ElasticSearchURL),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)
	failure.FailOnError(err, "Unable to initialise ElasticSearch")

	return client
}
