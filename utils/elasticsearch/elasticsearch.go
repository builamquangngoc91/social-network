package elasticsearch

import (
	"context"
	"encoding/json"
	log "social-network/utils/log"

	"github.com/olivere/elastic/v7"
)

type ElasticClient struct {
	esclient *elastic.Client
}

var l = log.New()

func NewElasticClient(urls []string) (*ElasticClient, error) {
	client, err := elastic.NewClient(elastic.SetURL(urls...),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	return &ElasticClient{
		esclient: client,
	}, err
}

func (e *ElasticClient) Index(ctx context.Context, index string, body interface{}) error {
	bodyJson, err := json.Marshal(body)
	if err != nil {
		l.Fatalf("Json error: %v", err)
	}

	_, err = e.esclient.Index().
		Index(index).
		BodyJson(string(bodyJson)).
		Do(ctx)

	return err
}

func (e *ElasticClient) Search(ctx context.Context, )
